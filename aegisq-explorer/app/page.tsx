"use client"

import { useEffect, useState } from "react"
import { getStatus } from "../lib/api"

export default function Page() {

  const [status, setStatus] = useState<any>(null)

  useEffect(() => {
    getStatus().then(setStatus)
  }, [])

  if (!status) return <div>Loading...</div>

  return (
    <div>

      <h1 className="text-3xl mb-6 font-bold">
        AegisQ Explorer
      </h1>

      <div className="bg-gray-900 p-6 rounded-lg border border-gray-800">

        <p>
          <span className="font-bold">Status:</span>{" "}
          <span className="text-green-400">{status.status}</span>
        </p>

        <p>
          <span className="font-bold">Latest Height:</span>{" "}
          {status.height}
        </p>

      </div>

    </div>
  )
}