package service

import (
	"errors"
	"project-capston/features/goverment"
	"project-capston/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	repo:=new(mocks.GovermentData)
	inputData:=goverment.Core{Type: "hospital",Name: "Rumah Sakit Wardiwaluyo"}

	t.Run("success create data",func(t *testing.T){
		repo.On("Insert",inputData).Return(nil).Once()
		srv:=New(repo)
		err:=srv.Create(inputData)
		assert.Nil(t,err)
		repo.AssertExpectations(t)
	})

	t.Run("fail validate",func(t *testing.T){
		srv:=New(repo)
		err:=srv.Create(goverment.Core{Type: "hospital"})
		assert.NotNil(t,err)
		repo.AssertExpectations(t)		
	})
}

func TestGetAll(t *testing.T) {
	repo:=new(mocks.GovermentData)
	returnData:=[]goverment.Core{{ID: uint(1), Type: "hospital",Name: "Rumah Sakit Wardiwaluyo"}}

	t.Run("success get all",func(t *testing.T){
		repo.On("SelectAll",int(1),int(1)).Return(returnData,nil).Once()
		srv:=New(repo)
		response,err:=srv.GetAll(int(1),int(1))
		assert.Equal(t,returnData,response)
		assert.Nil(t,err)
		repo.AssertExpectations(t)
	})

	t.Run("failed get all",func(t *testing.T){
		repo.On("SelectAll",int(1),int(1)).Return(nil,errors.New("fail get all goverment")).Once()
		srv:=New(repo)
		response,err:=srv.GetAll(int(1),int(1))
		assert.NotNil(t,err)
		assert.Nil(t,response)
		repo.AssertExpectations(t)
	})
}

func TestGetById(t *testing.T) {
	repo:=new(mocks.GovermentData)
	returnData:=goverment.Core{ID: uint(1), Type: "hospital",Name: "Rumah Sakit Wardiwaluyo"}

	t.Run("success get by id",func(t *testing.T){
		repo.On("Select",uint(1)).Return(returnData,nil).Once()
		srv:=New(repo)
		response,err:=srv.GetById(uint(1))
		assert.Equal(t,returnData,response)
		assert.Nil(t,err)
		repo.AssertExpectations(t)
	})
	t.Run("failed get by id",func(t *testing.T){
		repo.On("Select",uint(1)).Return(goverment.Core{},errors.New("error get by id")).Once()
		srv:=New(repo)
		response,err:=srv.GetById(uint(1))
		assert.Equal(t,goverment.Core{},response)
		assert.NotNil(t,err)
		repo.AssertExpectations(t)
	})
}

func TestEdit(t *testing.T) {
	repo:=new(mocks.GovermentData)
	inputData:=goverment.Core{ID: uint(1), Type: "hospital",Name: "Rumah Sakit Wardiwaluyo"}

	t.Run("success edit by id",func(t *testing.T){
		repo.On("Update",uint(1),inputData).Return(nil).Once()
		srv:=New(repo)
		err:=srv.EditById(uint(1),inputData)
		assert.Nil(t,err)
		repo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T){
	repo:=new(mocks.GovermentData)

	t.Run("success delete by id",func(t *testing.T){
		repo.On("Delete",uint(1)).Return(nil).Once()
		srv:=New(repo)
		err:=srv.DeleteById(uint(1))
		assert.Nil(t,err)
		repo.AssertExpectations(t)
	})	
}

func TestGetNearestLocation(t *testing.T){
	repo:=new(mocks.GovermentData)
	returnData:=[]goverment.Location{{ID: uint(1),Name: "Rumah Sakit Wardiwaluyo",Latitude: float64(9.56),Longitude: float64(7.88),Distance: float64(3)}}

	t.Run("success get nearest location",func(t *testing.T){
		repo.On("SelectNearestLocation",float64(9.56),float64(7.88),float64(10)).Return(returnData,nil).Once()
		srv:=New(repo)
		response,err:=srv.GetNearestLocation(float64(9.56),float64(7.88),float64(10))
		assert.Equal(t,returnData,response)
		assert.Nil(t,err)
		repo.AssertExpectations(t)
	})

	t.Run("failed get nearest location",func(t *testing.T){
		repo.On("SelectNearestLocation",float64(9.56),float64(7.88),float64(10)).Return(nil,errors.New("failed get nearest location")).Once()
		srv:=New(repo)
		response,err:=srv.GetNearestLocation(float64(9.56),float64(7.88),float64(10))
		assert.Nil(t,response)
		assert.NotNil(t,err)
		repo.AssertExpectations(t)
	})
}

func TestSelectCount(t *testing.T){
	repo:=new(mocks.GovermentData)
	returnData:=goverment.UnitCount{RumahSakit: int64(5),Pemadam: int64(3),Kepolisian: int64(3),SAR: int64(2),Dishub: int64(1)}

	t.Run("success get count",func(t *testing.T){
		repo.On("SelectCountUnit").Return(returnData,nil).Once()
		srv:=New(repo)
		response,err:=srv.GetCountUnit("superadmin")
		assert.Nil(t,err)
		assert.Equal(t,returnData,response)
		repo.AssertExpectations(t)
	})

	t.Run("failed get count",func(t *testing.T){
		repo.On("SelectCountUnit").Return(goverment.UnitCount{},errors.New("failed get count unit")).Once()
		srv:=New(repo)
		response,err:=srv.GetCountUnit("superadmin")
		assert.NotNil(t,err)
		assert.Equal(t,goverment.UnitCount{},response)
		repo.AssertExpectations(t)
	})
	t.Run("failed get count because not superadmin",func(t *testing.T){
		srv:=New(repo)
		response,err:=srv.GetCountUnit("admin")
		assert.NotNil(t,err)
		assert.Equal(t,goverment.UnitCount{},response)
		repo.AssertExpectations(t)
	})
}

