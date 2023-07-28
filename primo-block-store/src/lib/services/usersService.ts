import { createUser } from "$lib/domain/entities/user";
import type { IRecordServiceResponse } from "./serviceUtils";

export interface IUsersServiceResponse extends IRecordServiceResponse {
  "username": string,
  "email": string,
  "name": string,
  "avatar": string,
}

export function createUserByServiceResponse(data: IUsersServiceResponse) {
  return createUser(data)
}