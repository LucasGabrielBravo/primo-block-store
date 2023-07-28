import type { CategoryType } from "$lib/utils/constants";
import { createEntity, type ICreateEntityParams, type IEntity } from "./entity";
import { createUser, type ICreateUserParams, type IUser } from "./user";

export interface IBlock extends IEntity {
  user: IUser,
  name: string,
  symbol: Object,
  public: boolean,
  category: CategoryType
}

export interface ICreateBlockParams extends ICreateEntityParams, Omit<IBlock, keyof IEntity | 'user'> {
  user: IUser | ICreateUserParams
}

export function createBlock(params: ICreateBlockParams) {
  const { category, name, public: p, symbol } = params

  const user = createUser(params.user)

  return {
    ...createEntity(params),
    category, name, public: p, symbol, user
  }
}

