<script lang="ts">
    import { onMount } from "svelte";
    
    onMount(async () => {
        const ws = new WebSocket("ws://localhost:8080/buzzer/ws");
        ws.onopen = () => {
            console.log("ws (buzzer): Connected");
        }
        ws.onmessage = (event) => {
            const message = JSON.parse(event.data);
            console.log('Message received:', message);
            if (message.type === 'ack') {
                console.log('Ack received:', message.data);
            }
        }
        ws.onclose = () => {
            console.log("ws (buzzer): Disconnected");
        }
    })
        
</script>


<div role="alert" class="alert alert-warning my-2">
    <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" class="stroke-current shrink-0 w-6 h-6"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path></svg>
    <span>Veuillez presser le buzzer uniquement une fois par manche. Celui-ci se débloque automatiquement à la prochaine manche.</span>
</div>
<div class="flex items-center justify-center min-h-96">
<div class="card w-2/3 mx-auto">
        <div class="card-body items-center text-center bg-gray-200 rounded border-2 border-gray-400">
            <h2 class="card-title">Buzzer</h2>
            <button class="btn btn-outline btn-primary">Presser le buzzer</button>
        </div>
    </div>
</div>