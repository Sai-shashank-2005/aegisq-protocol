const API = "http://localhost:8080"

export async function getStatus() {
  const res = await fetch(`${API}/status`, {
    cache: "no-store"
  })
  return await res.json()
}

export async function getBlocks() {
  const res = await fetch(`${API}/blocks`, {
    cache: "no-store"
  })
  return await res.json()
}

export async function getBlock(height: number) {
  const res = await fetch(`${API}/block/${height}`, {
    cache: "no-store"
  })
  return await res.json()
}