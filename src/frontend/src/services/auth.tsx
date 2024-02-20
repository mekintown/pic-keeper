import axios from "@/libs/axios";
import {
  LoginCredentials,
  LoginResponse,
  NewUser,
  RegisterCustomerResponse,
} from "@/types";

const authBaseUrl = "/authen/v1";

const registerCustomer = async (newUser: NewUser) => {
  try {
    const response = await axios.post<RegisterCustomerResponse>(
      `${authBaseUrl}/register/customer`,
      newUser
    );
    return response.data;
  } catch (error) {
    throw error;
  }
};

const login = async (loginCredentials: LoginCredentials) => {
  try {
    const response = await axios.post<LoginResponse>(
      `${authBaseUrl}/login`,
      loginCredentials
    );
    return response.data;
  } catch (error) {
    throw error;
  }
};

const googleLogin = async () => {
  try {
    const response = await axios.post(`${authBaseUrl}/google/login`);
    return response.data;
  } catch (error) {
    throw error;
  }
};

const authService = { registerCustomer, login, googleLogin };

export default authService;
