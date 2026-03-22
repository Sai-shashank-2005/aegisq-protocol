// lib/api.ts

export const API = "http://localhost:8080"

// 🔥 GLOBAL FETCH OPTIONS
const options: RequestInit = {
  cache: "no-store"
}

// 🔥 SAFE FETCH WRAPPER (important for stability)
async function fetchJSON(url: string) {
  try {
    const res = await fetch(url, options)

    if (!res.ok) {
      throw new Error(`API Error: ${res.status}`)
    }

    return await res.json()
  } catch (err) {
    console.error("Fetch failed:", url, err)
    return null
  }
}

// -------------------- BASIC --------------------

export async function getStatus() {
  return fetchJSON(`${API}/status`)
}

export async function getBlocks() {
  return fetchJSON(`${API}/blocks`)
}

export async function getBlock(height: number) {
  return fetchJSON(`${API}/block/${height}`)
}

// 🔥 FIXED (this is your correct function name)
export async function getTx(height: number, index: number) {
  return fetchJSON(`${API}/tx/${height}/${index}`)
}

export async function getTxHash(hash: string) {
  return fetchJSON(`${API}/txhash/${hash}`)
}

// -------------------- SYSTEM --------------------

export async function getConsensus() {
  return fetchJSON(`${API}/consensus`)
}

export async function getLiveness() {
  return fetchJSON(`${API}/liveness`)
}