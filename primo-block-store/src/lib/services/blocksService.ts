import { createBlock, type IBlock } from "$lib/domain/entities/block";
import { pb } from "$lib/stores/pocketbase";
import type { CategoryType } from "$lib/utils/constants";
import { get } from "svelte/store";
import { createUserByServiceResponse, type IUsersServiceResponse } from "./usersService";

interface IBlockServiceResponse {
  "id": string,
  "collectionId": string,
  "collectionName": string,
  "created": string,
  "updated": string,
  "user": string,
  "name": string,
  "symbol": string,
  "public": boolean,
  "category": CategoryType,
  "expand": {
    "user": IUsersServiceResponse
  }
}

export function createBlockByServiceResponse(data: IBlockServiceResponse): IBlock {
  const user = createUserByServiceResponse(data.expand.user)

  return createBlock({ ...data, user })
}

export function createBlocksService() {
  const collection = get(pb).collection("blocks")

  async function getFullList() {
    try {
      const blocks = await collection.getFullList<IBlockServiceResponse>({ expand: "user" })
      console.log(blocks)

      return blocks.map(createBlockByServiceResponse)
    } catch (error) {
      console.log(error)

      return []
    }
  }

  return {
    getFullList
  }
}