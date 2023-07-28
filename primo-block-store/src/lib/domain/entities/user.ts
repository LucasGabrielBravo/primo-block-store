import { createEntity, type ICreateEntityParams, type IEntity } from "./entity"

export interface IUser extends IEntity{
  email: string,
  name: string,
  avatar: string
}

export interface ICreateUserParams extends ICreateEntityParams, Omit<IUser, keyof IEntity> {}

export function createUser(params: ICreateUserParams) {
  const { avatar, email, name } = params

  return { 
    ...createEntity(params),
    avatar, email, name, 
  }
}