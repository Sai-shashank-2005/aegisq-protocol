"use client"

import {useEffect,useState} from "react"
import {useParams} from "next/navigation"
import {getBlock} from "../../../lib/api"
import Link from "next/link"

export default function BlockPage(){

  const params = useParams()
  const height = Number(params.height)

  const [block,setBlock] = useState<any>(null)

  useEffect(()=>{
    getBlock(height).then(setBlock)
  },[height])

  if(!block) return <div>Loading...</div>

  return(

    <div>

      <h1 className="text-3xl mb-6">
        Block {block.Index}
      </h1>

      <div className="mb-4">
        Transactions: {block.Transactions.length}
      </div>

      <table className="w-full">

        <thead>
          <tr className="border-b border-gray-700">
            <th className="p-3 text-left">Tx Index</th>
            <th className="p-3 text-left">Sender</th>
          </tr>
        </thead>

        <tbody>

          {block.Transactions.slice(0,50).map((tx:any,i:number)=>(
            <tr key={i} className="border-b border-gray-800">

              <td className="p-3">

                <Link
                  className="text-blue-400"
                  href={`/tx/${height}/${i}`}
                >
                  {i}
                </Link>

              </td>

              <td className="p-3">
                {tx.sender_id}
              </td>

            </tr>
          ))}

        </tbody>

      </table>

    </div>
  )
}