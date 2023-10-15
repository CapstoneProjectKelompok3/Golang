package service

import (
	"errors"
	"project-capston/features/vehicles"
	"project-capston/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDelete(t *testing.T) {
	repo:=new(mocks.VehicleData)

	t.Run("success delete vehicle", func(t *testing.T){
		repo.On("Delete",uint(1)).Return(nil).Once()
		srv:=New(repo)
		err:=srv.Delete(uint(1))
		assert.Nil(t,err)
		repo.AssertExpectations(t)
	})

	t.Run("fail delete vehicle", func(t *testing.T){
		repo.On("Delete",uint(1)).Return(errors.New("failed delete vehicles")).Once()
		srv:=New(repo)
		err:=srv.Delete(uint(1))
		assert.NotNil(t,err)
		repo.AssertExpectations(t)
	})

}

func TestGetAll(t *testing.T) {
	repo:=new(mocks.VehicleData)
	returnData:=[]vehicles.VehicleEntity{{Id: uint(1),GovermentID: uint(1),Plate: "BE 8464",Status: true}}

	t.Run("success get all",func(t *testing.T){
		repo.On("SelectAll").Return(returnData,nil).Once()
		srv:=New(repo)
		response,err:=srv.GetAll()
		assert.Equal(t,returnData,response)
		assert.Nil(t,err)
		repo.AssertExpectations(t)
	})

	t.Run("failed get all",func(t *testing.T){
		repo.On("SelectAll").Return(nil,errors.New("error get all vehicles")).Once()
		srv:=New(repo)
		response,err:=srv.GetAll()
		assert.Nil(t,response)
		assert.NotNil(t,err)
		repo.AssertExpectations(t)		
	})
}

func TestGetById(t *testing.T) {
	repo:=new(mocks.VehicleData)
	returnData:=vehicles.VehicleEntity{Id: uint(1),GovermentID: uint(1),Plate: "BE 8464",Status: true}

	t.Run("success get by id",func( t *testing.T){
		repo.On("SelectById",uint(1)).Return(returnData,nil).Once()
		srv:=New(repo)
		response,err:=srv.GetById(uint(1))
		assert.Equal(t,returnData,response)
		assert.Nil(t,err)
		repo.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	repo:=new(mocks.VehicleData)
	inputData:=vehicles.VehicleEntity{Id: uint(1),GovermentID: uint(1),Plate: "BE 8464",Status: true}

	t.Run("success update vehicle",func(t *testing.T){
		repo.On("Update",inputData,uint(1)).Return(nil).Once()
		srv:=New(repo)
		err:=srv.Edit(inputData,uint(1),"admin")
		assert.Nil(t,err)
		repo.AssertExpectations(t)
	})

	t.Run("fail update vehicle because not admin",func(t *testing.T){
		srv:=New(repo)
		err:=srv.Edit(inputData,uint(1),"user")
		assert.NotNil(t,err)
		repo.AssertExpectations(t)
	})

	t.Run("failed update vehicle",func(t *testing.T){
		repo.On("Update",inputData,uint(1)).Return(errors.New("failed update vehicle")).Once()
		srv:=New(repo)
		err:=srv.Edit(inputData,uint(1),"admin")
		assert.NotNil(t,err)
		repo.AssertExpectations(t)
	})
}

func TestAdd(t *testing.T) {
	repo:=new(mocks.VehicleData)
	inputData:=vehicles.VehicleEntity{Id: uint(1),GovermentID: uint(1),Plate: "BE 8464",Status: true}
	
	t.Run("success add vehicle",func(t *testing.T){
		repo.On("Insert",inputData).Return(nil).Once()
		srv:=New(repo)
		err:=srv.Add(inputData,"admin")
		assert.Nil(t,err)
		repo.AssertExpectations(t)
	})

	t.Run("fail add vehicle because not admin",func(t *testing.T){
		srv:=New(repo)
		err:=srv.Add(inputData,"user")
		assert.NotNil(t,err)
		repo.AssertExpectations(t)
	})

	t.Run("fail validate",func(t *testing.T){
		srv:=New(repo)
		err:=srv.Add(vehicles.VehicleEntity{Id: uint(1),GovermentID: uint(1)},"admin")
		assert.NotNil(t,err)
		repo.AssertExpectations(t)
	})

	t.Run("fail add vehicle",func(t *testing.T){
		repo.On("Insert",inputData).Return(errors.New("error insert vehicle")).Once()
		srv:=New(repo)
		err:=srv.Add(inputData,"admin")
		assert.NotNil(t,err)
		repo.AssertExpectations(t)
	})
}