export default function StatCard({title,value}:{title:string,value:any}){

return(

<div className="bg-gray-900 border border-gray-800 rounded p-6">

<div className="text-gray-400 text-sm">
{title}
</div>

<div className="text-2xl mt-2">
{value}
</div>

</div>

)
}