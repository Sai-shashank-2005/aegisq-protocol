"use client"

import { useEffect, useState } from "react"
import { useParams } from "next/navigation"
import { getBlock } from "../../../lib/api"
import Link from "next/link"

export default function BlockPage() {

  const params = useParams()
  const height = Number(params.height)

  const [block, setBlock] = useState<any>(null)

  useEffect(() => {
    async function load() {
      const data = await getBlock(height)
      setBlock(data)
    }

    if (height) load()

  }, [height])

  if (!block) return <div>Loading block...</div>

  return (
    <div>

      <h1 className="text-3xl mb-6">
        Block {block.Index}
      </h1>

      <div className="bg-gray-900 border border-gray-800 p-6 rounded">

        <p><b>Height:</b> {block.Index}</p>
        <p><b>Transactions:</b> {block.Transactions.length}</p>

      </div>

      <h2 className="text-xl mt-6 mb-4">
        Transactions
      </h2>

      <table className="w-full">

        <thead className="bg-gray-800">
          <tr>
            <th className="p-3 text-left">Index</th>
            <th className="p-3 text-left">Sender</th>
            <th className="p-3 text-left">Hash</th>
          </tr>
        </thead>

        <tbody>

          {block.Transactions.slice(0,20).map((tx:any,i:number)=>(
            <tr key={i} className="border-t border-gray-800">

              <td className="p-3">
                <Link
                  href={`/tx/${block.Index}/${i}`}
                  className="text-blue-400"
                >
                  {i}
                </Link>
              </td>

              <td className="p-3">
                {tx.sender_id}
              </td>

              <td className="p-3 font-mono">
                {tx.data_hash?.slice(0,20)}...
              </td>

            </tr>
          ))}

        </tbody>

      </table>

    </div>
  )
}