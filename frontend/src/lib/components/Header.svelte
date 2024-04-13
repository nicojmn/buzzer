<script lang="ts">
  import "/src/app.css"
  import logo_akap from '$lib/images/logo_akap.png'
  import { onMount } from "svelte";

  let conn_state:any = {} as JSON

  onMount(async () => {
    const response =await fetch("/admin/isAdmin")
    conn_state = await response.json()
  })

</script>

<svelte:head>
  <meta charset="utf-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1" />
  <title>N'oubliez pas les paroles !</title>
  <link rel="icon" type="image/png" href="favicon.png">
</svelte:head>

{#if conn_state.isAdmin}
  <nav id="header-nav" class="flex flex-row bg-gray-500 p-6 ">
    <div class="flex items-center flex-shrink-0 text-white mr-6">
      <img src={logo_akap} alt="akap logo" class="fill-current h-10">
    </div>
    <div class="w-full flex flex-row space-x-2">
      <div>
        <a href="/admin/dashboard" class="btn-black">
          Dashboard
        </a>
      </div>
      <div>
        <a href="/buzzer" class="btn-black">
          Buzzers
        </a>
      </div>
      <div>
        <a href="/" class="btn-black">
          Autre ?
        </a>
      </div>
    </div>
  </nav>
{/if}



