package service

import (
	"fmt"
	"github.com/stretchr/testify/mock"
	"reflect"
	"testing"

	"github.com/rafimuhammad01/user-service/internal/domain"
	"github.com/rafimuhammad01/user-service/mocks"
)

func TestUser_GetAll(t *testing.T) {
	type fields struct {
		repository UserRepository
	}
	type mockRepo struct {
		output1 []*domain.User
		output2 []error
	}
	tests := []struct {
		name     string
		fields   fields
		mockRepo mockRepo
		want     []*domain.User
		want1    []error
	}{
		{
			name: "success",
			fields: fields{
				repository: mocks.NewUserRepository(t),
			},
			mockRepo: mockRepo{
				output1: []*domain.User{
					{
						ID:       "id",
						Name:     "test",
						Username: "test",
						Password: "test",
					},
				},
				output2: nil,
			},
			want: []*domain.User{
				{
					ID:       "id",
					Name:     "test",
					Username: "test",
					Password: "test",
				},
			},
			want1: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				repository: tt.fields.repository,
			}
			tt.fields.repository.(*mocks.UserRepository).On("GetAll").Return(tt.mockRepo.output1, tt.mockRepo.output2)
			got, got1 := u.GetAll()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("User.GetAll() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("User.GetAll() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestUser_Create(t *testing.T) {
	type fields struct {
		repository UserRepository
		hashAlgo   HashAlgo
	}
	type args struct {
		user domain.User
	}
	type mockRepoCreate struct {
		input  domain.User
		output []error
	}
	type mockRepoValidateUsername struct {
		input   string
		output1 *domain.User
		output2 []error
	}
	type mockHash struct {
		input1  string
		input2  int
		output1 string
		output2 []error
	}
	tests := []struct {
		name                        string
		fields                      fields
		args                        args
		mockRepoCreate              mockRepoCreate
		mockRepoValidateUsername    mockRepoValidateUsername
		mockHash                    mockHash
		useMockRepoCreate           bool
		useMockRepoValidateUsername bool
		useMockHash                 bool
		want                        []error
	}{
		{
			name: "success",
			fields: fields{
				repository: mocks.NewUserRepository(t),
				hashAlgo:   mocks.NewHashAlgo(t),
			},
			args: args{
				domain.User{
					ID:       "test",
					Name:     "test",
					Password: "test",
					Username: "test",
				},
			},
			useMockRepoCreate:           true,
			useMockRepoValidateUsername: true,
			useMockHash:                 true,
			mockRepoCreate: mockRepoCreate{
				input: domain.User{
					ID:       "test",
					Name:     "test",
					Password: "newTest",
					Username: "test",
				},
				output: nil,
			},
			mockRepoValidateUsername: mockRepoValidateUsername{
				input:   "test",
				output1: nil,
				output2: []error{fmt.Errorf("[%w]", domain.ErrUserNotFound)},
			},
			mockHash: mockHash{
				input1:  "test",
				input2:  14,
				output1: "newTest",
				output2: nil,
			},
			want: nil,
		},
		{
			name: "validate error",
			fields: fields{
				repository: mocks.NewUserRepository(t),
				hashAlgo:   mocks.NewHashAlgo(t),
			},
			args: args{
				domain.User{
					ID:       "test",
					Name:     "",
					Password: "testpassword",
					Username: "testusername",
				},
			},
			useMockRepoCreate:           false,
			useMockRepoValidateUsername: true,
			useMockHash:                 true,
			mockRepoValidateUsername: mockRepoValidateUsername{
				input: "testusername",
				output1: &domain.User{
					ID:       "test",
					Name:     "testName",
					Password: "testPassword",
					Username: "testUsername",
				},
				output2: nil,
			},
			mockHash: mockHash{
				input1:  "testpassword",
				input2:  14,
				output1: "newTest",
				output2: nil,
			},
			want: []error{fmt.Errorf("[%w] name should not be empty", domain.ErrUserInvalidData), fmt.Errorf("[%w] username testusername already exist", domain.ErrUserExist)},
		},
		{
			name: "validate error 2",
			fields: fields{
				repository: mocks.NewUserRepository(t),
				hashAlgo:   mocks.NewHashAlgo(t),
			},
			args: args{
				domain.User{
					ID:       "test",
					Name:     "testname",
					Password: "testpassword",
					Username: "testusername",
				},
			},
			useMockRepoCreate:           false,
			useMockRepoValidateUsername: true,
			useMockHash:                 true,
			mockRepoValidateUsername: mockRepoValidateUsername{
				input:   "testusername",
				output1: nil,
				output2: []error{fmt.Errorf("new error")},
			},
			mockHash: mockHash{
				input1:  "testpassword",
				input2:  14,
				output1: "",
				output2: []error{fmt.Errorf("new error")},
			},
			want: []error{fmt.Errorf("new error"), fmt.Errorf("new error")},
		},
		{
			name: "success",
			fields: fields{
				repository: mocks.NewUserRepository(t),
				hashAlgo:   mocks.NewHashAlgo(t),
			},
			args: args{
				domain.User{
					ID:       "test",
					Name:     "test",
					Password: "test",
					Username: "test",
				},
			},
			useMockRepoCreate:           true,
			useMockRepoValidateUsername: true,
			useMockHash:                 true,
			mockRepoCreate: mockRepoCreate{
				input: domain.User{
					ID:       "test",
					Name:     "test",
					Password: "newTest",
					Username: "test",
				},
				output: []error{fmt.Errorf("test new error")},
			},
			mockRepoValidateUsername: mockRepoValidateUsername{
				input:   "test",
				output1: nil,
				output2: []error{fmt.Errorf("[%w]", domain.ErrUserNotFound)},
			},
			mockHash: mockHash{
				input1:  "test",
				input2:  14,
				output1: "newTest",
				output2: nil,
			},
			want: []error{fmt.Errorf("test new error")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				repository: tt.fields.repository,
				hashAlgo:   tt.fields.hashAlgo,
			}

			if tt.useMockRepoCreate {
				tt.fields.repository.(*mocks.UserRepository).On("Create", mock.Anything).Return(tt.mockRepoCreate.output)
			}

			if tt.useMockRepoValidateUsername {
				tt.fields.repository.(*mocks.UserRepository).On("GetByUsername", tt.mockRepoValidateUsername.input).Return(tt.mockRepoValidateUsername.output1, tt.mockRepoValidateUsername.output2)
			}

			if tt.useMockHash {
				tt.fields.hashAlgo.(*mocks.HashAlgo).On("Hash", tt.mockHash.input1, tt.mockHash.input2).Return(tt.mockHash.output1, tt.mockHash.output2)
			}

			if got := u.Create(tt.args.user); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("User.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_Update(t *testing.T) {
	type fields struct {
		repository UserRepository
		hashAlgo   HashAlgo
	}
	type args struct {
		user domain.User
	}
	type mockGetByID struct {
		input   string
		output1 *domain.User
		output2 []error
	}
	type mockAlgoHash struct {
		input1  string
		input2  int
		output1 string
		output2 []error
	}
	type mockUpdate struct {
		input  domain.User
		output []error
	}
	type mockGetByUsername struct {
		input   string
		output1 *domain.User
		output2 []error
	}
	tests := []struct {
		name              string
		fields            fields
		args              args
		mockGetByID       mockGetByID
		mockAlgoHash      mockAlgoHash
		mockUpdate        mockUpdate
		mockGetByUsername mockGetByUsername
		isMockGetByID     bool
		isMockAlgoHash    bool
		isMockUpdate      bool
		isMockByUsername  bool
		want              []error
	}{
		{
			name: "another error when validating username",
			fields: fields{
				repository: mocks.NewUserRepository(t),
				hashAlgo:   mocks.NewHashAlgo(t),
			},
			args: args{
				domain.User{
					ID:       "id",
					Username: "testusername",
					Name:     "test name",
					Password: "testpassword",
				},
			},
			mockGetByID: mockGetByID{
				input: "id",
				output1: &domain.User{
					ID:       "id",
					Username: "testusername",
					Name:     "test name",
					Password: "testpassword",
				},
			},
			mockAlgoHash: mockAlgoHash{
				input1:  "testpassword",
				input2:  14,
				output1: "newtestpassword",
				output2: nil,
			},
			mockGetByUsername: mockGetByUsername{
				input:   "testusername",
				output1: nil,
				output2: []error{fmt.Errorf("new error")},
			},
			isMockByUsername: true,
			isMockGetByID:    true,
			isMockAlgoHash:   true,
			want:             []error{fmt.Errorf("new error")},
		},
		{
			name: "error when validating username",
			fields: fields{
				repository: mocks.NewUserRepository(t),
				hashAlgo:   mocks.NewHashAlgo(t),
			},
			args: args{
				user: domain.User{
					ID:       "id",
					Username: "testusername",
					Name:     "test name",
					Password: "testpassword",
				},
			},
			mockGetByID: mockGetByID{
				input: "id",
				output1: &domain.User{
					ID:       "id",
					Username: "testusername",
					Name:     "test name",
					Password: "testpassword",
				},
			},
			mockAlgoHash: mockAlgoHash{
				input1:  "testpassword",
				input2:  14,
				output1: "newtestpassword",
				output2: nil,
			},
			mockGetByUsername: mockGetByUsername{
				input:   "testusername",
				output1: &domain.User{},
				output2: nil,
			},
			isMockByUsername: true,
			isMockGetByID:    true,
			isMockAlgoHash:   true,
			want:             []error{fmt.Errorf("[%w] username testusername already exist", domain.ErrUserExist)},
		},
		{
			name: "success",
			fields: fields{
				repository: mocks.NewUserRepository(t),
				hashAlgo:   mocks.NewHashAlgo(t),
			},
			args: args{
				domain.User{
					ID:       "id",
					Username: "testusername",
					Name:     "test name",
					Password: "testpassword",
				},
			},
			mockGetByID: mockGetByID{
				input: "id",
				output1: &domain.User{
					ID:       "id",
					Username: "testusername",
					Name:     "test name",
					Password: "testpassword",
				},
			},
			mockAlgoHash: mockAlgoHash{
				input1:  "testpassword",
				input2:  14,
				output1: "newtestpassword",
				output2: nil,
			},
			mockUpdate: mockUpdate{
				input: domain.User{
					ID:       "id",
					Username: "testusername",
					Name:     "test name",
					Password: "newtestpassword",
				},
				output: nil,
			},
			mockGetByUsername: mockGetByUsername{
				input:   "testusername",
				output1: nil,
				output2: []error{fmt.Errorf("%w", domain.ErrUserNotFound)},
			},
			isMockByUsername: true,
			isMockGetByID:    true,
			isMockAlgoHash:   true,
			isMockUpdate:     true,
			want:             nil,
		},
		{
			name: "error from update repository",
			fields: fields{
				repository: mocks.NewUserRepository(t),
				hashAlgo:   mocks.NewHashAlgo(t),
			},
			args: args{
				domain.User{
					ID:       "id",
					Username: "testusername",
					Name:     "test name",
					Password: "testpassword",
				},
			},
			mockGetByID: mockGetByID{
				input: "id",
				output1: &domain.User{
					ID:       "id",
					Username: "testusername",
					Name:     "test name",
					Password: "testpassword",
				},
			},
			mockAlgoHash: mockAlgoHash{
				input1:  "testpassword",
				input2:  14,
				output1: "newtestpassword",
				output2: nil,
			},
			mockUpdate: mockUpdate{
				input: domain.User{
					ID:       "id",
					Username: "testusername",
					Name:     "test name",
					Password: "newtestpassword",
				},
				output: []error{fmt.Errorf("test error")},
			},
			mockGetByUsername: mockGetByUsername{
				input:   "testusername",
				output1: nil,
				output2: []error{fmt.Errorf("%w", domain.ErrUserNotFound)},
			},
			isMockByUsername: true,
			isMockGetByID:    true,
			isMockAlgoHash:   true,
			isMockUpdate:     true,
			want:             []error{fmt.Errorf("test error")},
		},
		{
			name: "error from repository",
			fields: fields{
				repository: mocks.NewUserRepository(t),
				hashAlgo:   mocks.NewHashAlgo(t),
			},
			args: args{
				domain.User{
					ID:       "id",
					Username: "testusername",
					Name:     "Test Name",
					Password: "testpassword",
				},
			},
			mockGetByID: mockGetByID{
				input:   "id",
				output1: nil,
				output2: []error{fmt.Errorf("test another error")},
			},
			mockAlgoHash: mockAlgoHash{
				input1:  "testpassword",
				input2:  14,
				output1: "newTestPassword",
				output2: nil,
			},
			mockGetByUsername: mockGetByUsername{
				input:   "testusername",
				output1: nil,
				output2: []error{fmt.Errorf("%w", domain.ErrUserNotFound)},
			},
			isMockByUsername: true,
			isMockGetByID:    true,
			isMockAlgoHash:   true,
			want:             []error{fmt.Errorf("test another error")},
		},
		{
			name: "error user not found",
			fields: fields{
				repository: mocks.NewUserRepository(t),
				hashAlgo:   mocks.NewHashAlgo(t),
			},
			args: args{
				domain.User{
					ID:       "id",
					Username: "testusername",
					Name:     "Test Name",
					Password: "testpassword",
				},
			},
			mockGetByID: mockGetByID{
				input:   "id",
				output1: nil,
				output2: []error{fmt.Errorf("[%w] id %s", domain.ErrUserNotFound, "id")},
			},
			mockAlgoHash: mockAlgoHash{
				input1:  "testpassword",
				input2:  14,
				output1: "testnewpassword",
				output2: nil,
			},
			mockGetByUsername: mockGetByUsername{
				input:   "testusername",
				output1: nil,
				output2: []error{fmt.Errorf("%w", domain.ErrUserNotFound)},
			},
			isMockByUsername: true,
			isMockGetByID:    true,
			isMockAlgoHash:   true,
			want: []error{
				fmt.Errorf("[%w] user with id %s not found", domain.ErrUserNotFound, "id"),
			},
		},
		{
			name: "error validate fields",
			fields: fields{
				repository: mocks.NewUserRepository(t),
				hashAlgo:   mocks.NewHashAlgo(t),
			},
			args: args{
				domain.User{
					ID:       "id",
					Username: "testusername",
					Name:     "Test Name",
				},
			},
			mockGetByID: mockGetByID{
				input: "id",
				output1: &domain.User{
					ID:       "name",
					Username: "testusername",
					Name:     "Test Name",
					Password: "testPassword",
				},
				output2: nil,
			},
			mockAlgoHash: mockAlgoHash{
				input1:  "",
				input2:  14,
				output1: "",
				output2: []error{fmt.Errorf("test error")},
			},
			mockGetByUsername: mockGetByUsername{
				input:   "testusername",
				output1: nil,
				output2: []error{fmt.Errorf("%w", domain.ErrUserNotFound)},
			},
			isMockByUsername: true,
			isMockGetByID:    true,
			isMockAlgoHash:   true,
			want: []error{
				fmt.Errorf("[%w] password should not be empty", domain.ErrUserInvalidData),
				fmt.Errorf("test error"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				repository: tt.fields.repository,
				hashAlgo:   tt.fields.hashAlgo,
			}

			if tt.isMockGetByID {
				tt.fields.repository.(*mocks.UserRepository).On("GetByID", tt.mockGetByID.input).Return(tt.mockGetByID.output1, tt.mockGetByID.output2)
			}

			if tt.isMockAlgoHash {
				tt.fields.hashAlgo.(*mocks.HashAlgo).On("Hash", tt.mockAlgoHash.input1, tt.mockAlgoHash.input2).Return(tt.mockAlgoHash.output1, tt.mockAlgoHash.output2)
			}

			if tt.isMockUpdate {
				tt.fields.repository.(*mocks.UserRepository).On("Update", mock.Anything).Return(tt.mockUpdate.output)
			}

			if tt.isMockByUsername {
				tt.fields.repository.(*mocks.UserRepository).On("GetByUsername", tt.mockGetByUsername.input).Return(tt.mockGetByUsername.output1, tt.mockGetByUsername.output2)
			}

			if got := u.Update(tt.args.user); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("User.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_GetByID(t *testing.T) {
	type fields struct {
		repository UserRepository
		hashAlgo   HashAlgo
	}
	type args struct {
		id string
	}
	type mockGetByID struct {
		input   string
		output1 *domain.User
		output2 []error
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		mockGetByID   mockGetByID
		isMockGetByID bool
		want          *domain.User
		want1         []error
	}{
		// TODO: Add test cases.
		{
			name: "success",
			fields: fields{
				repository: mocks.NewUserRepository(t),
				hashAlgo:   mocks.NewHashAlgo(t),
			},
			args: args{
				id: "testID",
			},
			mockGetByID: mockGetByID{
				input: "testID",
				output1: &domain.User{
					ID:       "testID",
					Name:     "test name",
					Username: "testusername",
					Password: "testpassword",
				},
				output2: nil,
			},
			isMockGetByID: true,
			want: &domain.User{
				ID:       "testID",
				Name:     "test name",
				Username: "testusername",
				Password: "testpassword",
			},
			want1: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := NewUser(tt.fields.repository, tt.fields.hashAlgo)

			if tt.isMockGetByID {
				tt.fields.repository.(*mocks.UserRepository).On("GetByID", tt.mockGetByID.input).Return(tt.mockGetByID.output1, tt.mockGetByID.output2)
			}

			got, got1 := u.GetByID(tt.args.id)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("User.GetByID() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("User.GetByID() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestUser_Delete(t *testing.T) {
	type fields struct {
		repository UserRepository
		hashAlgo   HashAlgo
	}
	type args struct {
		id string
	}
	type mockGetByID struct {
		input   string
		output1 *domain.User
		output2 []error
	}
	type mockDelete struct {
		input  string
		output []error
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		mockGetByID   mockGetByID
		mockDelete    mockDelete
		isMockGetByID bool
		isMockDelete  bool
		want          []error
	}{
		{
			name: "error from repo when delete user",
			fields: fields{
				mocks.NewUserRepository(t),
				mocks.NewHashAlgo(t),
			},
			args: args{
				id: "testID",
			},
			mockGetByID: mockGetByID{
				input: "testID",
				output1: &domain.User{
					ID:       "testID",
					Name:     "test name",
					Username: "testusername",
					Password: "testpassword",
				},
				output2: nil,
			},
			isMockGetByID: true,
			mockDelete: mockDelete{
				input: "testID",
				output: []error{
					fmt.Errorf("test error"),
				},
			},
			isMockDelete: true,
			want: []error{
				fmt.Errorf("test error"),
			},
		},
		{
			name: "error validate if user is exist",
			fields: fields{
				mocks.NewUserRepository(t),
				mocks.NewHashAlgo(t),
			},
			args: args{
				id: "testID",
			},
			mockGetByID: mockGetByID{
				input:   "testID",
				output1: nil,
				output2: []error{
					fmt.Errorf("[%w] id %s", domain.ErrUserNotFound, "testID"),
				},
			},
			isMockGetByID: true,
			want:          []error{fmt.Errorf("[%w] user with id %s not found", domain.ErrUserNotFound, "testID")},
		},
		{
			name: "another error from getbyid",
			fields: fields{
				mocks.NewUserRepository(t),
				mocks.NewHashAlgo(t),
			},
			args: args{
				id: "testID",
			},
			mockGetByID: mockGetByID{
				input:   "testID",
				output1: nil,
				output2: []error{fmt.Errorf("test error")},
			},
			isMockGetByID: true,
			want: []error{
				fmt.Errorf("test error"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				repository: tt.fields.repository,
				hashAlgo:   tt.fields.hashAlgo,
			}

			if tt.isMockGetByID {
				tt.fields.repository.(*mocks.UserRepository).On("GetByID", tt.mockGetByID.input).Return(tt.mockGetByID.output1, tt.mockGetByID.output2)
			}

			if tt.isMockDelete {
				tt.fields.repository.(*mocks.UserRepository).On("Delete", mock.Anything).Return(tt.mockDelete.output)
			}

			if got := u.Delete(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("User.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}
