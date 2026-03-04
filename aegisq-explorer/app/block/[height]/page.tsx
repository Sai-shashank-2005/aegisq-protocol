"use client"

import { useEffect, useState } from "react"
import { getBlock } from "../../../lib/api"

export default function BlockPage({ params }: any) {

  const [block, setBlock] = useState<any>(null)

  useEffect(() => {
    getBlock(params.height).then(setBlock)
  }, [])

  if (!block) return <div>Loading...</div>

  return (
    <div>

      <h1 className="text-3xl font-bold mb-6">
        Block {block.Index}
      </h1>

      <div className="bg-gray-900 p-6 rounded-lg border border-gray-800">

        <p>
          <span className="font-bold">Hash:</span>
          <span className="font-mono text-gray-400 ml-2">
            {block.Hash}
          </span>
        </p>

        <p>
          <span className="font-bold">Transactions:</span>{" "}
          {block.Transactions.length}
        </p>

      </div>

    </div>
  )
}