export default function MetricCard({title,value,color}:any){

return(

<div className="bg-gray-900 border border-gray-800 rounded-xl p-6">

<p className="text-gray-400 text-sm">
{title}
</p>

<p className={`text-2xl font-semibold mt-2 ${color}`}>
{value}
</p>

</div>

)

}