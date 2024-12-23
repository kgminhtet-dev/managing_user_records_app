export interface IUser {
  id: string;
  email: string;
  password: string;
  created_at: Date;
  updated_at: Date;
}

export interface ICreateUser {
  id: string;
  name: string;
  email: string;
  password: string;
}

export interface IUpdateUser {
  id: string;
  name: string;
  email: string;
}

export interface ILog {
  id: string;
  user_id: string;
  event: string;
  data: unknown;
  timestamp: Date;
}
