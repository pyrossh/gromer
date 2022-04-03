package components

func Header() string {
	return (`
		<div class="flex flex-row justify-center items-center w-full mb-20 font-bold text-xl text-gray-700 p-4">
			<div class="text-blue-700">
					<a href="https://pyros.sh"> pyros.sh </a>
			</div>
			<div class="flex flex-row flex-1 justify-end items-end p-2">
					<div class="border-b-2 border-white text-lg text-blue-700 mr-4">Examples:</div>
					<div class="border-b-2 border-white hover:border-red-700 mr-4">
							<a href="/"> Home </a>
					</div>
					<div class="border-b-2 border-white hover:border-red-700 mr-4">
							<a href="/clock"> Clock </a>
					</div>
					<div class="border-b-2 border-white hover:border-red-700 mr-4">
							<a href="/about"> About </a>
					</div>
					<div class="border-b-2 border-white hover:border-red-700 mr-4">
							<a href="/container"> Container </a>
					</div>
					<div class="border-b-2 border-white hover:border-red-700 mr-4">
							<a href="/panic"> Panic </a>
					</div>
			</div>
			{{ children }}
		</div>
	`)
}
