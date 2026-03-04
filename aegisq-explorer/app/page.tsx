"use client"

import { useEffect, useState } from "react"
import { getStatus } from "../lib/api"

export default function Dashboard() {

  const [status, setStatus] = useState<any>(null)

  useEffect(() => {
    getStatus().then(setStatus)
  }, [])

  if (!status) return <div>Loading...</div>

  return (
    <div>

      <h1 className="text-3xl mb-6">AegisQ Explorer</h1>

      <div className="bg-gray-900 border border-gray-800 p-6 rounded">

        <p>Status: <span className="text-green-400">{status.status}</span></p>

        <p>Latest Height: {status.height}</p>

      </div>

    </div>
  )
}