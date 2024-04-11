<script lang="ts">
    // WIP 
    // TODO : fetch variables from backend
    let buzzer_number:number = Math.round(Math.random() * 10)
    let state:string = "close"
    let podium:string[] = ["athjrjytkuytkty", "byjjytkeutk" ,"cyjtdytkdtykuktj(e"]
</script>

<div class="flex flex-row">
    <a class="card" href="/admin/buzzers">
        <span class="text-center"> Buzzer connectés : {buzzer_number}/13</span>
    </a>
    
    <a class="card" href="/admin/buzzers">
        <span class="text-center">État des buzzers : {state === "open" ? "ouverts" : "déverouillés"}</span><br>
        {#if state === "close"}
            <button class="btn-blue mt-1">Déverrouiller</button>
        {:else}
            <button class="btn-blue mt-1">Verrouiller</button>
        {/if}
    </a>

    <div class="card">
        {#each {length : 3} as _,i}
            <span class="overflow-hidden"> Team #{i + 1} : {podium[i]}</span><br>
        {/each}
    </div>
</div>


