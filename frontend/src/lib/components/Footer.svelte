<script lang="ts">
  import logo_akap from '$lib/images/logo_akap.png'
  import banner_llnux from '$lib/images/tux_llnux_title_banner.svg'
</script>

<footer>
    <img id="tux-footer" src={banner_llnux} alt="Louvain-li-Nux banner" width="10%">
    <p>Made with love by Louvain-li-Nux</p>
    <img id="akap-footer" src={logo_akap} alt="Akapella logo" width="5%">
</footer>
  
  
  <style>
    footer {
      background-color: #f0f0f0;
      padding: 1em;
      text-align: center;
      width: 100%;
      position: fixed;
      bottom: 0;
      display: flex;
      align-items: center;
  
    }
    
    #tux-footer {
      float: right;
      margin-right: 35%;
    }
  
    #akap-footer {
      float: left;
      margin-left: 35%;
    }
  
    p {
      font-size: 75%;
    }
  </style>