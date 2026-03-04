"use client"

import {useEffect,useState} from "react"
import {getBlocks} from "../../lib/api"
import Link from "next/link"

export default function BlocksPage(){

  const [blocks,setBlocks] = useState<any[]>([])

  useEffect(()=>{
    getBlocks().then(setBlocks)
  },[])

  return(

    <div>

      <h1 className="text-3xl mb-6">
        Blocks
      </h1>

      <table className="w-full">

        <thead>
          <tr className="border-b border-gray-700">
            <th className="p-3 text-left">Height</th>
            <th className="p-3 text-left">Hash</th>
            <th className="p-3 text-left">Transactions</th>
          </tr>
        </thead>

        <tbody>

          {blocks.map((b)=>(
            <tr key={b.height} className="border-b border-gray-800">

              <td className="p-3">

                <Link
                  href={`/block/${b.height}`}
                  className="text-blue-400"
                >
                  {b.height}
                </Link>

              </td>

              <td className="p-3">
                {b.hash.slice(0,20)}...
              </td>

              <td className="p-3">
                {b.txs}
              </td>

            </tr>
          ))}

        </tbody>

      </table>

    </div>
  )
}