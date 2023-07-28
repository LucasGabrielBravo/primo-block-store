export interface IEntity {
  id: string
  created: Date
  updated: Date
}

export interface ICreateEntityParams extends Omit<IEntity, "created" | "updated" | "id"> {
  id?: string
  created?: Date | string
  updated?: Date | string
}

export function createEntity(params: ICreateEntityParams): IEntity{
  const id = params.id || ""

  const created = (params.created) ? new Date(params.created) : new Date()
  const updated = (params.updated) ? new Date(params.updated) : new Date()

  return {
    id, created, updated
  }
}