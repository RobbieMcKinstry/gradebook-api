// This module exports redux action creators for the GradebookAPI API hosted at localhost:8081.
// Redux Thunk middleware or equivalent is required to use these action creators.
// It uses the axios javascript library for making the actual HTTP requests.
import * as actions from './GradebookAPIActions';
import axios from 'axios';

// Make me a new bug profile
// path is the request path, the format is "/api/bug-profile"
// data contains the action payload (request body)
// This function returns a promise which dispatches an error if the HTTP response is a 4xx or 5xx.
export const createBug-Profile = (path, data) => {
  return (dispatch) => {
    dispatch(actions.requestCreateBug-Profile());
    return axios({
      timeout: 20000,
      url: 'http://localhost:8081' + path,
      method: 'post',
      data: data,
      responseType: 'json'
    })
      .then((response) => {
        dispatch(actions.receiveCreateBug-ProfileSuccess(response.data, response.status));
      })
      .catch((response) => {
        dispatch(actions.receiveCreateBug-ProfileError(response.data, response.status));
      });
  };
};

// Create or login a user with GitHub OAuth
// path is the request path, the format is "/api/auth/gh"
// This function returns a promise which dispatches an error if the HTTP response is a 4xx or 5xx.
export const createAuth/Gh = (path) => {
  return (dispatch) => {
    dispatch(actions.requestCreateAuth/Gh());
    return axios({
      timeout: 20000,
      url: 'http://localhost:8081' + path,
      method: 'post',
      responseType: 'json'
    })
      .then((response) => {
        dispatch(actions.receiveCreateAuth/GhSuccess(response.data, response.status));
      })
      .catch((response) => {
        dispatch(actions.receiveCreateAuth/GhError(response.data, response.status));
      });
  };
};

// Sign up for the first time
// path is the request path, the format is "/api/user"
// data contains the action payload (request body)
// This function returns a promise which dispatches an error if the HTTP response is a 4xx or 5xx.
export const createUser = (path, data) => {
  return (dispatch) => {
    dispatch(actions.requestCreateUser());
    return axios({
      timeout: 20000,
      url: 'http://localhost:8081' + path,
      method: 'post',
      data: data,
      responseType: 'json'
    })
      .then((response) => {
        dispatch(actions.receiveCreateUserSuccess(response.data, response.status));
      })
      .catch((response) => {
        dispatch(actions.receiveCreateUserError(response.data, response.status));
      });
  };
};

// I don't want this bug profile anymore
// path is the request path, the format is "/api/bug-profile/:profileID"
// This function returns a promise which dispatches an error if the HTTP response is a 4xx or 5xx.
export const deleteBug-Profile = (path) => {
  return (dispatch) => {
    dispatch(actions.requestDeleteBug-Profile());
    return axios({
      timeout: 20000,
      url: 'http://localhost:8081' + path,
      method: 'delete',
      responseType: 'json'
    })
      .then((response) => {
        dispatch(actions.receiveDeleteBug-ProfileSuccess(response.data, response.status));
      })
      .catch((response) => {
        dispatch(actions.receiveDeleteBug-ProfileError(response.data, response.status));
      });
  };
};

// Show all of my bug profiles
// path is the request path, the format is "/api/bug-profile"
// This function returns a promise which dispatches an error if the HTTP response is a 4xx or 5xx.
export const listBug-Profile = (path) => {
  return (dispatch) => {
    dispatch(actions.requestListBug-Profile());
    return axios({
      timeout: 20000,
      url: 'http://localhost:8081' + path,
      method: 'get',
      responseType: 'json'
    })
      .then((response) => {
        dispatch(actions.receiveListBug-ProfileSuccess(response.data, response.status));
      })
      .catch((response) => {
        dispatch(actions.receiveListBug-ProfileError(response.data, response.status));
      });
  };
};

// Returns the GH token
// path is the request path, the format is "/api/GHtoken"
// This function returns a promise which dispatches an error if the HTTP response is a 4xx or 5xx.
export const readGHtoken = (path) => {
  return (dispatch) => {
    dispatch(actions.requestReadGHtoken());
    return axios({
      timeout: 20000,
      url: 'http://localhost:8081' + path,
      method: 'get',
      responseType: 'json'
    })
      .then((response) => {
        dispatch(actions.receiveReadGHtokenSuccess(response.data, response.status));
      })
      .catch((response) => {
        dispatch(actions.receiveReadGHtokenError(response.data, response.status));
      });
  };
};

// Get the detials about this particular account
// path is the request path, the format is "/api/user/:userID"
// This function returns a promise which dispatches an error if the HTTP response is a 4xx or 5xx.
export const readUser = (path) => {
  return (dispatch) => {
    dispatch(actions.requestReadUser());
    return axios({
      timeout: 20000,
      url: 'http://localhost:8081' + path,
      method: 'get',
      responseType: 'json'
    })
      .then((response) => {
        dispatch(actions.receiveReadUserSuccess(response.data, response.status));
      })
      .catch((response) => {
        dispatch(actions.receiveReadUserError(response.data, response.status));
      });
  };
};

// Show a single bug profile
// path is the request path, the format is "/api/bug-profile/:profileID"
// This function returns a promise which dispatches an error if the HTTP response is a 4xx or 5xx.
export const showBug-Profile = (path) => {
  return (dispatch) => {
    dispatch(actions.requestShowBug-Profile());
    return axios({
      timeout: 20000,
      url: 'http://localhost:8081' + path,
      method: 'get',
      responseType: 'json'
    })
      .then((response) => {
        dispatch(actions.receiveShowBug-ProfileSuccess(response.data, response.status));
      })
      .catch((response) => {
        dispatch(actions.receiveShowBug-ProfileError(response.data, response.status));
      });
  };
};

// Update my pre-existing bug profile
// path is the request path, the format is "/api/bug-profile/:profileID"
// data contains the action payload (request body)
// This function returns a promise which dispatches an error if the HTTP response is a 4xx or 5xx.
export const updateBug-Profile = (path, data) => {
  return (dispatch) => {
    dispatch(actions.requestUpdateBug-Profile());
    return axios({
      timeout: 20000,
      url: 'http://localhost:8081' + path,
      method: 'put',
      data: data,
      responseType: 'json'
    })
      .then((response) => {
        dispatch(actions.receiveUpdateBug-ProfileSuccess(response.data, response.status));
      })
      .catch((response) => {
        dispatch(actions.receiveUpdateBug-ProfileError(response.data, response.status));
      });
  };
};

// Adjust your account settings
// path is the request path, the format is "/api/user/:userID"
// data contains the action payload (request body)
// This function returns a promise which dispatches an error if the HTTP response is a 4xx or 5xx.
export const updateUser = (path, data) => {
  return (dispatch) => {
    dispatch(actions.requestUpdateUser());
    return axios({
      timeout: 20000,
      url: 'http://localhost:8081' + path,
      method: 'put',
      data: data,
      responseType: 'json'
    })
      .then((response) => {
        dispatch(actions.receiveUpdateUserSuccess(response.data, response.status));
      })
      .catch((response) => {
        dispatch(actions.receiveUpdateUserError(response.data, response.status));
      });
  };
};
