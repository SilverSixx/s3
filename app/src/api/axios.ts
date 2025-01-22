import axios from "axios";

const api = axios.create({
  baseURL: import.meta.env.VITE_API_HOST || "http://localhost:8080",
});

// api.interceptors.request.use(
//   (config) => {
//     const token =
//       import.meta.env.VITE_BEARER_TOKEN || "esRyFcyDpmcXumThcw86pQjmvkNtAz8N";

//     if (token) {
//       config.headers.Authorization = `Bearer ${token}`;
//     }

//     return config;
//   },
//   (error) => {
//     return Promise.reject(error);
//   }
// );

export default api;
