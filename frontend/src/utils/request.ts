import axios from "axios";
import type { AxiosResponse } from "axios";
import { ElMessage } from "element-plus";
import { useUserStore } from "@/pinia/modules/user";
import router from "@/router/index";

const service = axios.create({
  baseURL: "/api/v1",
});

// http request 拦截器
service.interceptors.request.use(
  (config: any) => {
    const userStore = useUserStore();
    config.headers = {
      token: userStore.token,
      "Content-Type": "application/json",
      ...config.headers,
    };

    return config;
  },
  (error: any) => {
    // 请求错误，这里可以用全局提示框进行提示
    ElMessage({
      message: error,
      type: "error",
    });

    return error;
  }
);

// 请求拦截器，内部根据返回值，重新组装，统一管理。
service.interceptors.response.use(
  (response: AxiosResponse) => {
    const userStore = useUserStore();

    if (response.headers["new-token"]) {
      userStore.setToken(response.headers["new-token"]);
    }

    if (response.data.code === 200 || response.headers.success === "true") {
      if (response.headers.msg) {
        response.data.msg = decodeURI(response.headers.msg);
      }

      return response.data;
    } else {
      ElMessage.error({
        message: response.data.msg || decodeURI(response.headers.msg),
        type: "error",
      });
      if (response.data.data && response.data.data.reload) {
        userStore.token = "";
        localStorage.clear();
        router.push({ name: "Login", replace: true });
      }

      return response.data.msg ? response.data : response;
    }
  },
  (error) => {
    switch (error.response.status) {
      case 404:
        ElMessage.error("查询错误，找不到要请求的资源！");
        break;
      default:
        if (error.response.data) {
          if (error.response.data.msg) {
            ElMessage.error(error.response.data.msg);
          } else {
            ElMessage.error("服务器错误，请重新刷新！");
          }
        }
        break;
    }
    return Promise.reject(error);
  }
);

export default service;