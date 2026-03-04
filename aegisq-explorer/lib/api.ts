export const API = "http://localhost:8080"

export async function getStatus() {
  const res = await fetch(`${API}/status`)
  return res.json()
}

export async function getBlocks() {
  const res = await fetch(`${API}/blocks`)
  return res.json()
}

export async function getBlock(height:number) {
  const res = await fetch(`${API}/block/${height}`)
  return res.json()
}

export async function getTx(height:number,index:number){
  const res = await fetch(`${API}/tx/${height}/${index}`)
  return res.json()
}

export async function getTxHash(hash:string){
  const res = await fetch(`${API}/txhash/${hash}`)
  return res.json()
}