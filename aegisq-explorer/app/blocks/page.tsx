"use client"

import { useEffect, useState } from "react"
import { getBlocks } from "../../lib/api"
import Link from "next/link"

export default function BlocksPage() {

  const [blocks, setBlocks] = useState<any[]>([])

  useEffect(() => {
    getBlocks().then((data) => {
      if (Array.isArray(data)) setBlocks(data)
    })
  }, [])

  return (
    <div>

      <h1 className="text-3xl mb-6">Latest Blocks</h1>

      <table className="w-full">

        <thead className="bg-gray-800">
          <tr>
            <th className="p-4 text-left">Height</th>
            <th className="p-4 text-left">Transactions</th>
            <th className="p-4 text-left">Hash</th>
          </tr>
        </thead>

        <tbody>

          {blocks.map((b) => (
            <tr key={b.height} className="border-t border-gray-800">

              <td className="p-4">
                <Link
                  href={`/block/${b.height}`}
                  className="text-blue-400"
                >
                  {b.height}
                </Link>
              </td>

              <td className="p-4">
                {b.txs}
              </td>

              <td className="p-4 font-mono">
                {b.hash.slice(0,20)}...
              </td>

            </tr>
          ))}

        </tbody>

      </table>

    </div>
  )
}