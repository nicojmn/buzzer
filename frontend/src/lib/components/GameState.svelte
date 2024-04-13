<script lang="ts">
    // WIP 
    // TODO : fetch variables from backend
    let buzzer_number:number = Math.round(Math.random() * 10)
    let state:string = "close"
    let podium:string[] = ["athjrjytkuytkty", "byjjytkeutk" ,"cyjtdytkdtykuktj(e"]

    export let teamsNumber:number
</script>

<div class="flex flex-row justify-around my-4">
    <div class="card w-1/4 bg-white border-2 border-gray-300 rounded-lg shadow">
        <div class="card-body">
            <span class="text-center"> Buzzer connectés : {buzzer_number}/{teamsNumber}</span>
        </div>
    </div>
    
    <div class="card w-1/4 bg-white border-2 border-gray-300 rounded-lg shadow">
        <div class="card-body">
            <span class="text-center">État des buzzers : {state === "open" ? "ouverts" : "déverouillés"}</span><br>
            {#if state === "close"}
                <button class="btn btn-outline btn-success mt-1">Déverrouiller</button>
            {:else}
                <button class="btn btn-outline btn-error mt-1">Verrouiller</button>
            {/if}
        </div>
    </div>

    <div class="card w-1/4 bg-white border-2 border-gray-300 rounded-lg shadow">
        <div class="card-body text-center">
            {#each {length : 3} as _,i}
                <span class="overflow-hidden"> Team #{i + 1} : {podium[i]}</span><br>
            {/each}
        </div>
    </div>
</div>


