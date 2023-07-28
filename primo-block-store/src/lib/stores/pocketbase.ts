import { PUBLIC_POCKETBASE_URL } from '$env/static/public'
import Pocketbase from 'pocketbase'
import { readable } from 'svelte/store'

function createPocketbase() {
  const pb = new Pocketbase(PUBLIC_POCKETBASE_URL)  

  return readable(pb)
}

export const pb = createPocketbase()