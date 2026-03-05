export default function Hero({title,subtitle}:any){

return(

<div className="bg-gradient-to-r from-gray-900 to-gray-800 border border-gray-800 rounded-xl p-8">

<h1 className="text-3xl font-bold text-white">
{title}
</h1>

<p className="text-gray-400 mt-2">
{subtitle}
</p>

</div>

)

}