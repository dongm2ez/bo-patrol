import { get, post, put, del } from '@/utils/request'
import type { SpaceLocation, SpaceTree, CreateSpaceRequest, MapElement, CreateMapElementRequest, OptionItem } from '@/types'

export const getSpaces = (params?: { type?: string; floor?: number }) => {
  return get<{ list: SpaceLocation[], total: number }>('/spaces', params)
}

export const getSpaceTree = () => {
  return get<SpaceTree[]>('/spaces/tree')
}

export const getSpace = (id: number) => {
  return get<SpaceLocation>(`/spaces/${id}`)
}

export const createSpace = (data: CreateSpaceRequest) => {
  return post<SpaceLocation>('/spaces', data)
}

export const updateSpace = (id: number, data: CreateSpaceRequest) => {
  return put<SpaceLocation>(`/spaces/${id}`, data)
}

export const deleteSpace = (id: number) => {
  return del(`/spaces/${id}`)
}

export const getSpaceTypes = () => {
  return get<OptionItem[]>('/public/space-types')
}

export const getMapElementsByLocation = (locationId: number) => {
  return get<{ list: MapElement[], total: number }>(`/map-elements/location/${locationId}`)
}

export const createMapElement = (data: CreateMapElementRequest) => {
  return post<MapElement>('/map-elements', data)
}

export const updateMapElement = (id: number, data: CreateMapElementRequest) => {
  return put<MapElement>(`/map-elements/${id}`, data)
}

export const deleteMapElement = (id: number) => {
  return del(`/map-elements/${id}`)
}

export const batchUpdateMapElements = (data: MapElement[]) => {
  return put('/map-elements/batch', data)
}

export const getElementTypes = () => {
  return get<OptionItem[]>('/public/element-types')
}
