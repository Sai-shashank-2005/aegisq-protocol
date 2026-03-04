"use client"

import {useEffect,useState} from "react"
import {getStatus,getBlocks} from "../lib/api"
import Link from "next/link"

export default function Home(){

  const [status,setStatus] = useState<any>(null)
  const [blocks,setBlocks] = useState<any[]>([])

  useEffect(()=>{

    getStatus().then(setStatus)
    getBlocks().then(setBlocks)

  },[])

  return(

    <div>

      <h1 className="text-3xl mb-6">
        AegisQ Network
      </h1>

      {status && (
        <div className="mb-6">
          Latest Height: {status.height}
        </div>
      )}

      <h2 className="text-xl mb-4">
        Latest Blocks
      </h2>

      <table className="w-full">

        <thead>
          <tr className="border-b border-gray-700">
            <th className="p-3 text-left">Height</th>
            <th className="p-3 text-left">Transactions</th>
          </tr>
        </thead>

        <tbody>

          {blocks.slice(0,10).map((b)=>(
            <tr key={b.height} className="border-b border-gray-800">

              <td className="p-3">
                <Link
                  className="text-blue-400"
                  href={`/block/${b.height}`}
                >
                  {b.height}
                </Link>
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