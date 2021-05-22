package server

//func TestSessionServer(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	mockSesUsecase := mock_auth.NewMockUsecase(ctrl)
//
//	sessionsURL := "localhost:9997"
//	go StartSessionsGRPCServer(mockSesUsecase, sessionsURL)
//	grpcConn, err := grpc.Dial(sessionsURL, grpc.WithInsecure())
//	assert.Equal(t, err, nil)
//
//	session := &models.Sessions{
//		UserID:     "1",
//		Expiration: 86400,
//	}
//	expectedResult := models.Result{
//		ID:     "1",
//		Status: "OK",
//	}
//
//	sessionsClient := client.NewSessionsClient(grpcConn)
//	mockSesUsecase.EXPECT().CheckSession("1").Return(session, nil)
//	res, err := sessionsClient.Check(context.Background(), "1")
//	if err != nil {
//		t.Error(err)
//	}
//	if res != expectedResult {
//		t.Errorf("expected: %v\n got: %v", expectedResult, res)
//	}
//
//	mockSesUsecase.EXPECT().CreateSession("1").Return(session, nil)
//	res, err = sessionsClient.Create(context.Background(), "1")
//	if err != nil {
//		t.Error(err)
//	}
//	if res != expectedResult {
//		t.Errorf("expected: %v\n got: %v", expectedResult, res)
//	}
//
//	mockSesUsecase.EXPECT().DeleteSession("1").Return(nil)
//	res, err = sessionsClient.Delete(context.Background(), "1")
//	if err != nil {
//		t.Error(err)
//	}
//	if res.ID != "-1" {
//		t.Errorf("expected: %v\n got: %v", expectedResult, res)
//	}
//}
