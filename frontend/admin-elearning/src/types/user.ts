type Role = "ADMIN" | "TEACHER" | "STUDENT";

export interface User {
  numberNo?: number;
  userId: number;
  username: string;
  email: string;
  role: Role;
  password?: string;
  totalRows : number
}

export type Paging = {
  pageNumber : number;
  rowsPerPage : number;
}

export type UserInfo = {
  userId: number;
  username: string;
  email: string;
  role: Role;
  parentId: number;
  createdAt: string;
  createdBy: number;
  status: string;
  updatedAt: string;
  updatedBy: number;
};
