import Link from "next/link"

export default function BlockCard({block}:any){

return(

<Link href={`/block/${block.height}`}>

<div className="bg-blue-950 border border-blue-900 p-5 rounded hover:bg-blue-900">

<div className="text-lg font-semibold">
Block {block.height}
</div>

<div className="text-sm text-gray-400">
Transactions: {block.txs}
</div>

<div className="text-xs mt-2 text-gray-500 truncate">
{block.hash}
</div>

</div>

</Link>

)
}