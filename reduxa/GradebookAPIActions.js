// This module exports redux actions for the GradebookAPI API hosted at localhost:8081.
import * as types from './GradebookAPIActionTypes';

export const requestCreateBug-Profile = () => ({
  type: types.REQ_CREATE_BUG-PROFILE
});
export const receiveCreateBug-ProfileSuccess = (json, status) => ({
  type: types.RECV_CREATE_BUG-PROFILE_SUCCESS,
  data: json,
  message: false,
  status: status
});
export const receiveCreateBug-ProfileError = (json, status) => ({
  type: types.RECV_CREATE_BUG-PROFILE_ERROR,
  data: false,
  message: json,
  status: status
});
export const requestCreateAuth/Gh = () => ({
  type: types.REQ_CREATE_AUTH/GH
});
export const receiveCreateAuth/GhSuccess = (json, status) => ({
  type: types.RECV_CREATE_AUTH/GH_SUCCESS,
  data: json,
  message: false,
  status: status
});
export const receiveCreateAuth/GhError = (json, status) => ({
  type: types.RECV_CREATE_AUTH/GH_ERROR,
  data: false,
  message: json,
  status: status
});
export const requestCreateUser = () => ({
  type: types.REQ_CREATE_USER
});
export const receiveCreateUserSuccess = (json, status) => ({
  type: types.RECV_CREATE_USER_SUCCESS,
  data: json,
  message: false,
  status: status
});
export const receiveCreateUserError = (json, status) => ({
  type: types.RECV_CREATE_USER_ERROR,
  data: false,
  message: json,
  status: status
});
export const requestDeleteBug-Profile = () => ({
  type: types.REQ_DELETE_BUG-PROFILE
});
export const receiveDeleteBug-ProfileSuccess = (json, status) => ({
  type: types.RECV_DELETE_BUG-PROFILE_SUCCESS,
  data: json,
  message: false,
  status: status
});
export const receiveDeleteBug-ProfileError = (json, status) => ({
  type: types.RECV_DELETE_BUG-PROFILE_ERROR,
  data: false,
  message: json,
  status: status
});
export const requestListBug-Profile = () => ({
  type: types.REQ_LIST_BUG-PROFILE
});
export const receiveListBug-ProfileSuccess = (json, status) => ({
  type: types.RECV_LIST_BUG-PROFILE_SUCCESS,
  data: json,
  message: false,
  status: status
});
export const receiveListBug-ProfileError = (json, status) => ({
  type: types.RECV_LIST_BUG-PROFILE_ERROR,
  data: false,
  message: json,
  status: status
});
export const requestLoginUser = () => ({
  type: types.REQ_LOGIN_USER
});
export const receiveLoginUserSuccess = (json, status) => ({
  type: types.RECV_LOGIN_USER_SUCCESS,
  data: json,
  message: false,
  status: status
});
export const receiveLoginUserError = (json, status) => ({
  type: types.RECV_LOGIN_USER_ERROR,
  data: false,
  message: json,
  status: status
});
export const requestReadGHtoken = () => ({
  type: types.REQ_READ_GHTOKEN
});
export const receiveReadGHtokenSuccess = (json, status) => ({
  type: types.RECV_READ_GHTOKEN_SUCCESS,
  data: json,
  message: false,
  status: status
});
export const receiveReadGHtokenError = (json, status) => ({
  type: types.RECV_READ_GHTOKEN_ERROR,
  data: false,
  message: json,
  status: status
});
export const requestReadUser = () => ({
  type: types.REQ_READ_USER
});
export const receiveReadUserSuccess = (json, status) => ({
  type: types.RECV_READ_USER_SUCCESS,
  data: json,
  message: false,
  status: status
});
export const receiveReadUserError = (json, status) => ({
  type: types.RECV_READ_USER_ERROR,
  data: false,
  message: json,
  status: status
});
export const requestShowBug-Profile = () => ({
  type: types.REQ_SHOW_BUG-PROFILE
});
export const receiveShowBug-ProfileSuccess = (json, status) => ({
  type: types.RECV_SHOW_BUG-PROFILE_SUCCESS,
  data: json,
  message: false,
  status: status
});
export const receiveShowBug-ProfileError = (json, status) => ({
  type: types.RECV_SHOW_BUG-PROFILE_ERROR,
  data: false,
  message: json,
  status: status
});
export const requestUpdateBug-Profile = () => ({
  type: types.REQ_UPDATE_BUG-PROFILE
});
export const receiveUpdateBug-ProfileSuccess = (json, status) => ({
  type: types.RECV_UPDATE_BUG-PROFILE_SUCCESS,
  data: json,
  message: false,
  status: status
});
export const receiveUpdateBug-ProfileError = (json, status) => ({
  type: types.RECV_UPDATE_BUG-PROFILE_ERROR,
  data: false,
  message: json,
  status: status
});
export const requestUpdateUser = () => ({
  type: types.REQ_UPDATE_USER
});
export const receiveUpdateUserSuccess = (json, status) => ({
  type: types.RECV_UPDATE_USER_SUCCESS,
  data: json,
  message: false,
  status: status
});
export const receiveUpdateUserError = (json, status) => ({
  type: types.RECV_UPDATE_USER_ERROR,
  data: false,
  message: json,
  status: status
});
