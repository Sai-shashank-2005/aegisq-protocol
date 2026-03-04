"use client"

import { useEffect, useState } from "react"
import Link from "next/link"
import { getBlocks } from "../../lib/api"

export default function BlocksPage() {

  const [blocks, setBlocks] = useState<any[]>([])
  const [loading, setLoading] = useState(true)

  useEffect(() => {

    async function load() {
      const data = await getBlocks()

      console.log("API RESPONSE:", data)

      if (Array.isArray(data)) {
        setBlocks(data)
      }

      setLoading(false)
    }

    load()

  }, [])

  return (
    <div>

      <h1 className="text-3xl font-bold mb-6">
        Latest Blocks
      </h1>

      <div className="bg-gray-900 border border-gray-800 rounded-lg overflow-hidden">

        <table className="w-full text-sm">

          <thead className="bg-gray-800 text-white">
            <tr>
              <th className="p-4 text-left">Height</th>
              <th className="p-4 text-left">Transactions</th>
              <th className="p-4 text-left">Hash</th>
            </tr>
          </thead>

          <tbody>

            {loading && (
              <tr>
                <td colSpan={3} className="p-6 text-center">
                  Loading blocks...
                </td>
              </tr>
            )}

            {!loading && blocks.length === 0 && (
              <tr>
                <td colSpan={3} className="p-6 text-center">
                  No blocks found
                </td>
              </tr>
            )}

            {blocks.map((b) => (
              <tr
                key={b.height}
                className="border-t border-gray-800 hover:bg-gray-800"
              >
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

                <td className="p-4 font-mono text-gray-400">
                  {b.hash.slice(0,20)}...
                </td>
              </tr>
            ))}

          </tbody>

        </table>

      </div>

    </div>
  )
}