"use client"

import {useEffect,useState} from "react"
import {useParams} from "next/navigation"
import {getBlock} from "../../../lib/api"
import Link from "next/link"

export default function BlockPage(){

  const params = useParams()
  const height = Number(params.height)

  const [block,setBlock] = useState<any>(null)
  const [page,setPage] = useState(0)

  const PAGE_SIZE = 25

  useEffect(()=>{
    getBlock(height).then(setBlock)
  },[height])

  if(!block) return <div>Loading...</div>

  const start = page * PAGE_SIZE
  const end = start + PAGE_SIZE

  const txs = block.Transactions.slice(start,end)

  return(

    <div>

      <div className="text-gray-400 mb-4">
        <Link href="/">Home</Link> / 
        <Link href="/blocks"> Blocks</Link> / 
        Block {block.Index}
      </div>

      <h1 className="text-3xl mb-6">
        Block {block.Index}
      </h1>

      <div className="grid grid-cols-3 gap-4 mb-6">

        <div className="bg-gray-900 p-4 rounded">
          <div className="text-gray-400">Height</div>
          <div className="text-xl">{block.Index}</div>
        </div>

        <div className="bg-gray-900 p-4 rounded">
          <div className="text-gray-400">Transactions</div>
          <div className="text-xl">{block.Transactions.length}</div>
        </div>

        <div className="bg-gray-900 p-4 rounded">
          <div className="text-gray-400">View</div>
          <div className="text-xl">{block.View}</div>
        </div>

      </div>

      <table className="w-full">

        <thead className="border-b border-gray-700">
          <tr>
            <th className="p-3 text-left">Tx Index</th>
            <th className="p-3 text-left">Sender</th>
          </tr>
        </thead>

        <tbody>

          {txs.map((tx:any,i:number)=>{

            const index = start + i

            return(

              <tr key={index} className="border-b border-gray-800">

                <td className="p-3">

                  <Link
                    href={`/tx/${height}/${index}`}
                    className="text-blue-400"
                  >
                    {index}
                  </Link>

                </td>

                <td className="p-3">
                  {tx.sender_id}
                </td>

              </tr>

            )
          })}

        </tbody>

      </table>

      <div className="flex gap-4 mt-6">

        <button
          onClick={()=>setPage(Math.max(page-1,0))}
          className="bg-gray-800 px-4 py-2 rounded"
        >
          Previous
        </button>

        <button
          onClick={()=>setPage(page+1)}
          className="bg-gray-800 px-4 py-2 rounded"
        >
          Next
        </button>

      </div>

    </div>
  )
}