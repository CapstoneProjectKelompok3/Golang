package service

import (
	"errors"
	"project-capston/features/emergency"
	"project-capston/helper"
	"project-capston/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInclose(t *testing.T) {
	repo:=new(mocks.EmergencyData)
	
	t.Run("success is close",func(t *testing.T){
		repo.On("IncloseEmergency",uint(1)).Return(nil).Once()
		srv:=New(repo)
		err :=srv.IncloseEmergency(uint(1))
		assert.NoError(t,err)
		repo.AssertExpectations(t)
	})

	t.Run("failed is close",func(t *testing.T){
		repo.On("IncloseEmergency",uint(1)).Return(errors.New("error is close")).Once()
		srv:=New(repo)
		err :=srv.IncloseEmergency(uint(1))
		assert.Error(t,err)
		repo.AssertExpectations(t)		
	})

}
func TestSumEmergency(t *testing.T) {
	repo:=new(mocks.EmergencyData)
	t.Run("success sum emergency",func(t *testing.T){
		repo.On("SumEmergency").Return(int64(1),nil).Once()
		srv:=New(repo)
		count,err:=srv.SumEmergency()
		assert.Nil(t,err)
		assert.Equal(t,int64(1),count)
		repo.AssertExpectations(t)
	})
	t.Run("fail sum emergency",func(t *testing.T){
		repo.On("SumEmergency").Return(int64(0),errors.New("failed sum emergency")).Once()
		srv:=New(repo)
		count,err:=srv.SumEmergency()
		assert.NotNil(t,err)
		assert.Equal(t,int64(0),count)
		repo.AssertExpectations(t)		
	})
}

func TestActionGmail(t *testing.T) {
	repo:=new(mocks.EmergencyData)
	t.Run("success action gmail",func(t *testing.T){
		repo.On("ActionGmail","message").Return(nil).Once()
		srv:=New(repo)
		err:=srv.ActionGmail("message")
		assert.Nil(t,err)
		repo.AssertExpectations(t)
	})

	t.Run("fail action gmail",func(t *testing.T){
		repo.On("ActionGmail","message").Return(errors.New("failed action gmail")).Once()
		srv:=New(repo)
		err:=srv.ActionGmail("message")
		assert.NotNil(t,err)
		repo.AssertExpectations(t)		
	})
	
}

func TestGetAll(t *testing.T) {
	repo:=new(mocks.EmergencyData)
	returnData:=[]emergency.EmergencyEntity{{Id: uint(1),Name: "laporan 1",CallerID: uint(1),ReceiverID: uint(2),Latitude: float64(9.7),Longitude: float64(7.4)}}
	param:=emergency.QueryParams{Page: int(1),ItemsPerPage: int(2),IsClassDashboard: true}
	t.Run("success get all",func(t *testing.T){
		repo.On("SelectAll",param,"token",uint(1),"user").Return(int64(23),returnData,nil).Once()
		srv:=New(repo)
		_,response,err:=srv.GetAll(param,"token",uint(1),"user")
		assert.Nil(t,err)
		assert.Equal(t,returnData,response)
		repo.AssertExpectations(t)
	})
	t.Run("fail get all",func(t *testing.T){
		repo.On("SelectAll",param,"token",uint(1),"user").Return(int64(0),nil,errors.New("error get all")).Once()
		srv:=New(repo)
		_,response,err:=srv.GetAll(param,"token",uint(1),"user")
		assert.NotNil(t,err)
		assert.Nil(t,response)
		repo.AssertExpectations(t)
	})

	t.Run("count == 0",func(t *testing.T){
		repo.On("SelectAll",param,"token",uint(1),"user").Return(int64(0),returnData,nil).Once()
		srv:=New(repo)
		bol,response,err:=srv.GetAll(param,"token",uint(1),"user")
		assert.Nil(t,err)
		assert.NotNil(t,response)
		assert.False(t,bol)
		repo.AssertExpectations(t)
	})
	t.Run("IsClassDashboard is true",func(t *testing.T){
		repo.On("SelectAll",param,"token",uint(1),"user").Return(int64(7),returnData,nil).Once()
		srv:=New(repo)
		bol,response,err:=srv.GetAll(param,"token",uint(1),"user")
		assert.Nil(t,err)
		assert.NotNil(t,response)
		assert.False(t,bol)
		repo.AssertExpectations(t)
	})

	t.Run("data nil",func(t *testing.T){
		repo.On("SelectAll",param,"token",uint(1),"user").Return(int64(7),nil,nil).Once()
		srv:=New(repo)
		bol,response,err:=srv.GetAll(param,"token",uint(1),"user")
		assert.Nil(t,err)
		assert.Nil(t,response)
		assert.False(t,bol)
		repo.AssertExpectations(t)
	})
}

func TestGetById(t *testing.T) {
	repo:=new(mocks.EmergencyData)
	returnData:=emergency.EmergencyEntity{Id: uint(1),Name: "laporan 1",CallerID: uint(1),ReceiverID: uint(2),Latitude: float64(9.7),Longitude: float64(7.4)}

	t.Run("success get by id",func(t *testing.T){
		repo.On("SelectById",uint(1),"token").Return(returnData,nil).Once()
		srv:=New(repo)
		response,err:=srv.GetById(uint(1),"token")
		assert.Nil(t,err)
		assert.Equal(t,returnData,response)
		repo.AssertExpectations(t)
	})

	t.Run("failed get by id",func(t *testing.T){
		repo.On("SelectById",uint(1),"token").Return(emergency.EmergencyEntity{},errors.New("failed get by id")).Once()
		srv:=New(repo)
		response,err:=srv.GetById(uint(1),"token")
		assert.NotNil(t,err)
		assert.Equal(t,emergency.EmergencyEntity{},response)
		repo.AssertExpectations(t)
	})
	
}

func TestEdit(t *testing.T) {
	repo:=new(mocks.EmergencyData)
	inputData:=emergency.EmergencyEntity{Id: uint(1),Name: "laporan 1",CallerID: uint(1),ReceiverID: uint(2),Latitude: float64(9.7),Longitude: float64(7.4)}

	t.Run("success edit emergency",func(t *testing.T){
		repo.On("Update",inputData,uint(1)).Return(nil).Once()
		srv:=New(repo)
		err:=srv.Edit(inputData,uint(1),"admin",uint(1))
		assert.Nil(t,err)
		repo.AssertExpectations(t)
	})
	t.Run("failed edit emergency by user",func(t *testing.T){
		srv:=New(repo)
		err:=srv.Edit(inputData,uint(1),"user",uint(1))
		assert.NotNil(t,err)
		repo.AssertExpectations(t)
	})

	t.Run("failed edit emergency",func(t *testing.T){
		repo.On("Update",inputData,uint(1)).Return(errors.New("error edit emergency")).Once()
		srv:=New(repo)
		err:=srv.Edit(inputData,uint(1),"admin",uint(1))
		assert.NotNil(t,err)
		repo.AssertExpectations(t)
	})

}

func TestDelete(t *testing.T) {
	repo:=new(mocks.EmergencyData)
	t.Run("success delete",func(t *testing.T){
		repo.On("Delete",uint(1)).Return(nil).Once()
		srv:=New(repo)
		err:=srv.Delete(uint(1))
		assert.Nil(t,err)
		repo.AssertExpectations(t)
	})

	t.Run("fail delete",func(t *testing.T){
		repo.On("Delete",uint(1)).Return(errors.New("error delete emergency")).Once()
		srv:=New(repo)
		err:=srv.Delete(uint(1))
		assert.NotNil(t,err)
		repo.AssertExpectations(t)
	})
}

func TestAdd(t *testing.T) {
	repo:=new(mocks.EmergencyData)
	inputData:=emergency.EmergencyEntity{Name:"Kasus 1", CallerID: uint(1),ReceiverID: uint(2),Latitude: float64(9.7),Longitude: float64(7.4)}
	returnUser:=emergency.UserEntity{ID: int(1),Name: "Rajih",Email: "rajihhh28@gmail.com",Level: "user",EmailActive: true}
	notif:=helper.MessageGomailE{EmailReceiver: "rajihhh28@gmail.com",Sucject: "Kasus 1",Content: "Kasus terbaru yang harus sedang ditangani, semoga user dapat tenang dan menunggu notifikasi selanjutnya",Name: "Rajih",Email: "rajihhh28@gmail.com"}

	t.Run("success add emergency",func (t *testing.T){
		repo.On("Insert",inputData).Return(uint(1),nil).Once()
		repo.On("Update",inputData,uint(1)).Return(nil).Once()
		repo.On("SelectUser","1","token").Return(returnUser,nil).Once()
		repo.On("SendNotification",notif).Return("email terkirim",nil).Once()
		srv:=New(repo)
		id,err:= srv.Add(inputData,"token")
		assert.Equal(t,uint(1),id)
		assert.Nil(t,err)
		repo.AssertExpectations(t)
	})
	t.Run("fail validate add emergency",func (t *testing.T){
		srv:=New(repo)
		id,err:= srv.Add(emergency.EmergencyEntity{},"token")
		assert.Equal(t,uint(0),id)
		assert.NotNil(t,err)
		repo.AssertExpectations(t)
	})

	t.Run("failed insert emergency",func (t *testing.T){
		repo.On("Insert",inputData).Return(uint(0),errors.New("error insert emergency")).Once()
		srv:=New(repo)
		id,err:= srv.Add(inputData,"token")
		assert.Equal(t,uint(0),id)
		assert.NotNil(t,err)
		repo.AssertExpectations(t)
	})

	t.Run("failed create name emergency",func (t *testing.T){
		repo.On("Insert",inputData).Return(uint(1),nil).Once()
		repo.On("Update",inputData,uint(1)).Return(errors.New("error update emergency")).Once()
		srv:=New(repo)
		id,err:= srv.Add(inputData,"token")
		assert.Equal(t,uint(0),id)
		assert.NotNil(t,err)
		repo.AssertExpectations(t)
	})
	t.Run("failed select user emergency",func (t *testing.T){
		repo.On("Insert",inputData).Return(uint(1),nil).Once()
		repo.On("Update",inputData,uint(1)).Return(nil).Once()
		repo.On("SelectUser","1","token").Return(emergency.UserEntity{},errors.New("error get user")).Once()
		srv:=New(repo)
		id,err:= srv.Add(inputData,"token")
		assert.Equal(t,uint(0),id)
		assert.NotNil(t,err)
		repo.AssertExpectations(t)
	})

	t.Run("failed send email emergency",func (t *testing.T){
		repo.On("Insert",inputData).Return(uint(1),nil).Once()
		repo.On("Update",inputData,uint(1)).Return(nil).Once()
		repo.On("SelectUser","1","token").Return(returnUser,nil).Once()
		repo.On("SendNotification",notif).Return("email tidak terkirim",errors.New("failed send email")).Once()
		srv:=New(repo)
		id,err:= srv.Add(inputData,"token")
		assert.Equal(t,uint(0),id)
		assert.NotNil(t,err)
		repo.AssertExpectations(t)
	})
}