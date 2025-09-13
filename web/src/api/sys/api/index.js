import request from '@/utils/request';

// @Tags api
// @Summary Get the role list by paging
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "Get the role list by paging"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"Get Success"}"
// @Router /api/getApiList [post]
// {
//  page     int
//	pageSize int
// }
export async function getApiList(data) {
  const res = await request.post('/api/getApiList', data);
  if (res.data.code === 0) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// @Tags Api
// @Summary Creating a basic API
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.CreateApiParams true "Creating a basic API"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"Get Success"}"
// @Router /api/createApi [post]
export async function createApi(data) {
  const res = await request.post('/api/createApi', data);
  if (res.data.code === 0) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// @Tags menu
// @Summary Get the menu by id
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.GetById true "Get the menu by id"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"Get Success"}"
// @Router /menu/getApiById [post]
export async function getApiById(data) {
  const res = await request.post('/api/getApiById', data);
  if (res.data.code === 0) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// @Tags Api
// @Summary Update API
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.CreateApiParams true "Update API"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"Get Success"}"
// @Router /api/updateApi [post]
export async function updateApi(data) {
  const res = await request.post('/api/updateApi', data);
  if (res.data.code === 0) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// @Tags Api
// @Summary Set API authorization
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.CreateApiParams true "Set API authorization"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"Get Success"}"
// @Router /api/setAuthApi [post]
export async function setAuthApi(data) {
  const res = await request.post('/api/setAuthApi', data);
  if (res.data.code === 0) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// @Tags Api
// @Summary Get all APIs without paging
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} json "{"success":true,"data":{},"msg":"Get Success"}"
// @Router /api/getAllApis [post]
export async function getAllApis(data) {
  const res = await request.post('/api/getAllApis', data);
  if (res.data.code === 0) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// @Tags Api
// @Summary Delete specified API
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dbModel.Api true "Delete API"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"Get Success"}"
// @Router /api/deleteApi [post]
export async function deleteApi(data) {
  const res = await request.post('/api/deleteApi', data);
  if (res.data.code === 0) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// @Tags SysApi
// @Summary Delete selected APIs
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Get Success"}"
// @Router /api/deleteApisByIds [delete]
export async function deleteApisByIds(data) {
  const res = await request.post('/api/deleteApisByIds', data);
  if (res.data.code === 0) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// FreshCasbin
// @Tags      SysApi
// @Summary   Refresh Casbin cache
// @accept    application/json
// @Produce   application/json
// @Success   200   {object}  response.Response{msg=string}  "Get Success"
// @Router    /api/freshCasbin [get]
export async function freshCasbin(params) {
  const res = await request.get('/api/freshCasbin', { params });
  if (res.data.code === 0) {
    return res.data.msg;
  }
  return Promise.reject(new Error(res.data.msg));
}

export async function syncApi(params) {
  const res = await request.get('/api/syncApi', { params });
  if (res.data.code === 0) {
    return res.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

export async function getApiGroups(params) {
  const res = await request.get('/api/getApiGroups', { params });
  if (res.data.code === 0) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

export async function ignoreApi(data) {
  const res = await request.post('/api/ignoreApi', data);
  if (res.data.code === 0) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

export async function enterSyncApi(data) {
  const res = await request.post('/api/enterSyncApi', data);
  if (res.data.code === 0) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}
