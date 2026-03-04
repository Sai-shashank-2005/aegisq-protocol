const API = "http://localhost:8080"

export async function getStatus() {
  const res = await fetch(`${API}/status`, { cache: "no-store" })
  return await res.json()
}

export async function getBlocks() {
  const res = await fetch(`${API}/blocks`, { cache: "no-store" })

  try {
    return await res.json()
  } catch {
    return []
  }
}

export async function getBlock(height: number) {
  const res = await fetch(`${API}/block/${height}`, { cache: "no-store" })
  return await res.json()
}

export async function getTransaction(height: number, index: number) {
  const res = await fetch(`${API}/tx/${height}/${index}`, { cache: "no-store" })
  return await res.json()
}