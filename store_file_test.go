package versions

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stevepartridge/go/file"
)

func Test_Unit_Store_File_Init(t *testing.T) {

	makeNewFileStore(t)

	store.Migrations()

	if !file.Exists(fileStoreFilename) {
		t.Errorf("File %s does not exist", fileStoreFilename)
	}

	if localDb.getNextId() != 1 {
		t.Error("Next ID should be 1")
	}

	if err := localDb.load(); err != nil {
		t.Error("Load store should not be nil")
	}

}

func Test_Unit_Store_File_NewStoreWithFile(t *testing.T) {

	removeFileStore(t)

	store, err := NewStore("file")
	if err != nil {
		t.Errorf("Error creating new file store should be nil. %s",
			fileStoreFilename)
	}

	store.Migrations()

	if !file.Exists(fileStoreFilename) {
		t.Errorf("File %s does not exist", fileStoreFilename)
	}

}

func Test_Unit_Store_File_NewStoreWithEmptyString(t *testing.T) {

	removeFileStore(t)

	store, err := NewStore("")
	if err != nil {
		t.Errorf("Error creating new file store should be nil. %s",
			fileStoreFilename)
	}

	store.Migrations()

	if !file.Exists(fileStoreFilename) {
		t.Errorf("File %s does not exist", fileStoreFilename)
	}

	if store.Type() != "file" {
		t.Errorf("Store type should be file, but is %s", store.Type())
	}

}

func Test_Unit_Store_File_MigrationsCreate(t *testing.T) {

	store, err := newFileStore()
	if err != nil {
		t.Errorf("Error creating new file store should be nil. %s",
			fileStoreFilename)
	}

	store.Migrations()

	if !file.Exists(fileStoreFilename) {
		t.Errorf("File %s does not exist", fileStoreFilename)
	}

	if localDb.getNextId() != 1 {
		t.Error("Next ID should be 1")
	}

	if err := localDb.load(); err != nil {
		t.Error("Load store should not be nil")
	}

}

func Test_Unit_Store_FileCorruptFile(t *testing.T) {

	store, err := newFileStore()
	if err != nil {
		t.Errorf("Error creating new file store should be nil. %s",
			fileStoreFilename)
	}

	store.Migrations()

	if !file.Exists(fileStoreFilename) {
		t.Errorf("File %s does not exist", fileStoreFilename)
	}

	data, err := ioutil.ReadFile(fileStoreFilename)
	if err != nil {
		t.Error("Error loading file store file.")
	}

	bad := string(data) + "{,"

	err = ioutil.WriteFile(fileStoreFilename, []byte(bad), os.ModeAppend)
	if err != nil {
		t.Error("Error loading file store file.")
	}

	if err := localDb.load(); err == nil {
		t.Errorf("Load store err should not be nil but saw %s", err.Error())
	}

}

func Test_Unit_Store_File_CreateSuccess(t *testing.T) {

	makeNewFileStore(t)

	err := store.CreateApplication(&testApplication1)
	if err != nil {
		t.Errorf("Error creating new application should be nil. %s", err.Error())
	}

	version := testVersion1
	version.Id = 0

	err = store.Create(&version)
	if err != nil {
		t.Errorf("Error creating new version should be nil. %s", err.Error())
	}

}

func Test_Unit_Store_File_CreateFail_Invalid(t *testing.T) {

	makeNewFileStore(t)

	version := testVersion1

	err := store.Create(&version)
	if err == nil {
		t.Error("Error creating new version should be error.")
	}

	if err != ErrVersionIdNotEmpty {
		t.Errorf("Create error should be %s but was %s",
			ErrVersionIdNotEmpty.Error(),
			err.Error(),
		)
	}

	version.Id = 0
	version.ApplicationId = ""

	err = store.Create(&version)
	if err == nil {
		t.Error("Error creating new version should be error.")
	}

	if err != ErrVersionMissingApplicationId {
		t.Errorf("Create error should be %s but was %s",
			ErrVersionMissingApplicationId.Error(),
			err.Error(),
		)
	}

}

func Test_Unit_Store_File_CreateFail_Exists(t *testing.T) {

	makeNewFileStore(t)

	err := store.CreateApplication(&testApplication1)
	if err != nil {
		t.Errorf("Error creating new application should be nil. %s", err.Error())
	}

	version := testVersion1
	version.Id = 0

	err = store.Create(&version)
	if err != nil {
		t.Errorf("Error creating new version should be nil. %s", err)
	}

	if version.Id == 0 {
		t.Error("Version Id should not be 0")
	}

	err = store.Create(&version)
	if err != nil {
		if err != ErrVersionIdNotEmpty {
			t.Errorf("Error should be %s, but is %s",
				ErrVersionIdNotEmpty.Error(),
				err.Error(),
			)
		}
	}

	version.Id = 0

	err = store.Create(&version)
	if err != nil {
		if err != ErrVersionExists {
			t.Errorf("Error should be %s, but is %s",
				ErrVersionExists.Error(),
				err.Error(),
			)
		}
	}

}

func Test_Unit_Store_File_GetByIdSuccess(t *testing.T) {

	makeNewFileStore(t)

	err := store.CreateApplication(&testApplication1)
	if err != nil {
		t.Errorf("Error creating new application should be nil. %s", err.Error())
	}

	version := testVersion1
	version.Id = 0

	err = store.Create(&version)
	if err != nil {
		t.Errorf("Error creating new version should be nil. %s", err)
	}

	ver, err := store.GetById(version.Id)
	if err != nil {
		t.Errorf("Error should be nil but gave %s", err.Error())
	}

	if ver.Id != version.Id {
		t.Error("Version ids do not match")
	}

}

func Test_Unit_Store_File_GetLatest_Success(t *testing.T) {

	makeNewFileStore(t)

	err := store.CreateApplication(&testApplication1)
	if err != nil {
		t.Errorf("Error creating new application should be nil. %s", err.Error())
	}

	version := testVersion1
	version.Id = 0

	err = store.Create(&version)
	if err != nil {
		t.Errorf("Error creating new version should be nil. %s", err)
	}

	ver, err := store.GetLatest(version.ApplicationId)
	if err != nil {
		t.Errorf("Error should be nil but gave %s", err.Error())
	}

	if ver.Id != version.Id {
		t.Error("Version ids do not match")
	}

}

func Test_Unit_Store_File_GetByApplication_Success(t *testing.T) {

	makeNewFileStore(t)

	err := store.CreateApplication(&testApplication1)
	if err != nil {
		t.Errorf("Error creating new application should be nil. %s", err.Error())
	}

	version := testVersion1
	version.Id = 0

	err = store.Create(&version)
	if err != nil {
		t.Errorf("Error creating new version should be nil. %s", err)
	}

	version2 := testVersion2
	version2.Id = 0

	err = store.Create(&version2)
	if err != nil {
		t.Errorf("Error creating new version should be nil. %s", err)
	}

	vers, err := store.GetByApplication(version.ApplicationId)
	if err != nil {
		t.Errorf("Error should be nil but gave %s", err.Error())
	}

	if len(vers) != 2 {
		t.Errorf("Versions should be 2 but found %d", len(vers))
	}

	vers, err = store.GetByApplication(version.ApplicationId, 1)
	if err != nil {
		t.Errorf("Error should be nil but gave %s", err.Error())
	}

	if len(vers) != 1 {
		t.Error("Found %d version(s) but should have been 1", len(vers))
	}

}

func Test_Unit_Store_File_GetLatestStable_Success(t *testing.T) {

	makeNewFileStore(t)

	err := store.CreateApplication(&testApplication1)
	if err != nil {
		t.Errorf("Error creating new application should be nil. %s", err.Error())
	}

	version := testVersion1
	version.Id = 0

	err = store.Create(&version)
	if err != nil {
		t.Errorf("Error creating new version should be nil. %s", err)
	}

	ver, err := store.GetLatest(version.ApplicationId, 1)
	if err != nil {
		t.Errorf("Error should be nil but gave %s", err.Error())
	}

	if ver.Id != version.Id {
		t.Error("Version ids do not match")
	}

	if !ver.Stable {
		t.Error("Version stable should be true")
	}

}

func Test_Unit_Store_File_UpdateSuccess(t *testing.T) {

	makeNewFileStore(t)

	err := store.CreateApplication(&testApplication1)
	if err != nil {
		t.Errorf("Error creating new application should be nil. %s", err.Error())
	}

	version := testVersion1
	version.Id = 0

	err = store.Create(&version)
	if err != nil {
		t.Errorf("Error creating new version should be nil. %s", err)
	}

	version.ApplicationId = "999"

	err = store.Update(&version)
	if err != nil {
		t.Errorf("Error should be nil but gave %s", err.Error())
	}

	if version.ApplicationId != "999" {
		t.Error("Version application id doesn't match")
	}

}

func Test_Unit_Store_File_UpdateMissingAndInvalidId(t *testing.T) {

	makeNewFileStore(t)

	err := store.CreateApplication(&testApplication1)
	if err != nil {
		t.Errorf("Error creating new application should be nil. %s", err.Error())
	}

	version := testVersion1
	version.Id = 0

	err = store.Create(&version)
	if err != nil {
		t.Errorf("Error creating new version should be nil. %s", err)
	}

	version.Id = 0
	version.ApplicationId = "999"

	err = store.Update(&version)
	if err == nil {
		t.Errorf("Error should not be nil")
	}

	if err != ErrVersionMissingId {
		t.Error("Version application id doesn't match")
	}

	version.Id = 99
	err = store.Update(&version)
	if err == nil {
		t.Errorf("Error should not be nil")
	}

	if err != ErrVersionNotFound {
		t.Errorf("Version err should be %s but saw: %s",
			ErrVersionNotFound,
			err,
		)
	}

}

func Test_Unit_Store_File_UpdateMissingComponents(t *testing.T) {

	makeNewFileStore(t)

	err := store.CreateApplication(&testApplication1)
	if err != nil {
		t.Errorf("Error creating new application should be nil. %s", err.Error())
	}

	version := testVersion1
	version.Id = 0

	err = store.Create(&version)
	if err != nil {
		t.Errorf("Error creating new version should be nil. %s", err)
	}

	version.Id = 0
	version.ApplicationId = "999"
	version.Major = ""
	version.Minor = ""
	version.Revision = ""

	err = store.Update(&version)
	if err == nil {
		t.Errorf("Error should not be nil")
	}

	if err != ErrVersionMissingComponents {
		t.Errorf("Error should be %s but was %s",
			ErrVersionMissingComponents.Error(),
			err.Error(),
		)
	}

}

func Test_Unit_Store_File_DeleteSuccess(t *testing.T) {

	makeNewFileStore(t)

	err := store.CreateApplication(&testApplication1)
	if err != nil {
		t.Errorf("Error creating new application should be nil. %s", err.Error())
	}

	version := testVersion1
	version.Id = 0

	err = store.Create(&version)
	if err != nil {
		t.Errorf("Error creating new version should be nil. %s", err)
	}

	err = store.Delete(version.Id)
	if err != nil {
		t.Errorf("Error should be nil but gave %s", err.Error())
	}

	ver, err := store.GetById(version.Id)

	if ver.Id != 0 {
		t.Error("Version id should be 0")
	}

}

func Test_Unit_Store_File_DeleteNotExists(t *testing.T) {

	makeNewFileStore(t)

	err := store.Delete(99)
	if err == nil {
		t.Errorf("Error should not be nil")
	}

	if err != ErrVersionNotFound {
		t.Errorf("Error should have been %s but was %s", ErrVersionNotFound, err.Error())
	}

}

func Test_Unit_Store_File_CreateGetNextId_Success(t *testing.T) {

	makeNewFileStore(t)

	err := store.CreateApplication(&testApplication1)
	if err != nil {
		t.Errorf("Error creating new application should be nil. %s", err.Error())
	}

	version := testVersion1
	version.Id = 0

	err = store.Create(&version)
	if err != nil {
		t.Error("Error creating new version should be error.")
	}

	if localDb.getNextId() != 2 {
		t.Errorf("Next ID should be 2 not %d", localDb.getNextId())
	}

}

func Test_Unit_Store_File_CreateApplication_Success(t *testing.T) {

	makeNewFileStore(t)

	err := store.CreateApplication(&testApplication1)
	if err != nil {
		t.Errorf("Error creating new application should be nil. %s", err.Error())
	}

}

func Test_Unit_Store_File_UpdateApplication_Success(t *testing.T) {

	makeNewFileStore(t)

	err := store.CreateApplication(&testApplication1)
	if err != nil {
		t.Errorf("Error updating application should be nil. %s", err.Error())
	}

	newAppName := "App Name Updated"

	testApplication1.Name = newAppName
	err = store.UpdateApplication(&testApplication1)
	if err != nil {
		t.Errorf("Error updated application should be nil. %s", err.Error())
	}

	if testApplication1.Name != newAppName {
		t.Errorf("Name should have been %s but is %s",
			newAppName,
			testApplication1.Name)
	}

}

func Test_Unit_Store_File_CreateApplicationExists(t *testing.T) {

	makeNewFileStore(t)

	err := store.CreateApplication(&testApplication1)
	if err != nil {
		t.Errorf("Error creating new application should be nil. %s", err.Error())
	}

	err = store.CreateApplication(&testApplication1)
	if err == nil {
		t.Error("Error creating new existing application should not be nil")
	}

	if err != ErrApplicationExists {
		t.Errorf("Error should be %s but received %s",
			ErrApplicationExists.Error(),
			err.Error(),
		)
	}

}

func Test_Unit_Store_File_CreateApplicationInavlidId(t *testing.T) {

	makeNewFileStore(t)

	app := testApplication1

	app.Id = ""

	err := store.CreateApplication(&app)
	if err == nil {
		t.Error("Error creating new application should not be nil")
	}

	if err != ErrApplicationMissingId {
		t.Errorf("Error should be %s but received %s",
			ErrApplicationMissingId.Error(),
			err.Error(),
		)
	}

	app.Id = "Invalid-%"

	err = store.CreateApplication(&app)
	if err == nil {
		t.Error("Error creating new application should not be nil")
	}

	if err.Error() != ErrApplicationIdInvalidCharacter("%").Error() {
		t.Errorf("Error should be %s but received %s",
			ErrApplicationIdInvalidCharacter("%").Error(),
			err.Error(),
		)
	}

	app.Id = "Invalid ID"
	err = store.CreateApplication(&app)
	if err == nil {
		t.Error("Error creating new application should not be nil")
	}

	if err.Error() != ErrApplicationIdInvalidCharacter("space").Error() {
		t.Errorf("Error should be %s but received %s",
			ErrApplicationIdInvalidCharacter("space").Error(),
			err.Error(),
		)
	}

}

func Test_Unit_Store_File_CreateApplicationInavlidName(t *testing.T) {

	makeNewFileStore(t)

	app := testApplication1

	app.Name = ""

	err := store.CreateApplication(&app)
	if err == nil {
		t.Error("Error creating new application should not be nil")
	}

	if err != ErrApplicationMissingName {
		t.Errorf("Error should be %s but received %s",
			ErrApplicationMissingName.Error(),
			err.Error(),
		)
	}

}

func Test_Unit_Store_File_GetApplicationNotExists(t *testing.T) {

	makeNewFileStore(t)

	_, err := store.GetApplicationById("99")
	if err == nil {
		t.Error("Error creating new existing application should not be nil")
	}

	if err != ErrApplicationNotFound {
		t.Errorf("Error should be %s but received %s",
			ErrApplicationNotFound.Error(),
			err.Error(),
		)
	}

}

func Test_Unit_Store_File_DeleteApplication_Success(t *testing.T) {

	makeNewFileStore(t)

	err := store.CreateApplication(&testApplication1)
	if err != nil {
		t.Errorf("Error creating new application should be nil. %s", err.Error())
	}

	err = store.DeleteApplication(testApplication1.Id)
	if err != nil {
		t.Errorf("Delete should not result in error %s", err.Error())
	}

	if store.ApplicationExists(&testApplication1) {
		t.Error("Application should not exist")
	}
}

func Test_Unit_Store_File_DeleteApplication_NotFound(t *testing.T) {

	makeNewFileStore(t)

	err := store.DeleteApplication("999")
	if err == nil {
		t.Error("Error deleting application should not be nil")
	}

	if err != ErrApplicationNotFound {
		t.Errorf("Error should be %s but received %s",
			ErrApplicationNotFound.Error(),
			err.Error(),
		)
	}
}

func Test_Unit_Store_File_GetApplications_Success(t *testing.T) {

	makeNewFileStore(t)

	err := store.CreateApplication(&testApplication1)
	if err != nil {
		t.Errorf("Error creating new application should be nil. %s", err.Error())
	}

	err = store.CreateApplication(&testApplication2)
	if err != nil {
		t.Errorf("Error creating new application should be nil. %s", err.Error())
	}

	apps, err := store.GetApplications()
	if err != nil {
		t.Errorf("Error get applications should be nil. %s", err.Error())
	}

	if len(apps) != 2 {
		t.Errorf("Should have found 2 apps but got %d", len(apps))
	}

	apps, err = store.GetApplications(1)
	if err != nil {
		t.Errorf("Error get applications should be nil. %s", err.Error())
	}

	if len(apps) != 1 {
		t.Errorf("Should have found 1 app but got %d", len(apps))
	}

}
