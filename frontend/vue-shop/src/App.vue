<script>
    import { RouterLink, RouterView } from 'vue-router'
</script>

<template>
    <link href="https://fonts.googleapis.com/icon?family=Material+Icons" rel="stylesheet">
    <div> 
      <nav>
        <div class="nav-wrapper lime accent-3">
          <RouterLink to="/"><a href="#" class="brand-logo black-text text-black">RacketShop</a></RouterLink>
          <ul id="nav-mobile" class="right hide-on-med-and-down">
            <li><RouterLink to="/cart">
              <i class="material-icons black-text text-black">shopping_cart</i>
            </RouterLink></li>
            <li><RouterLink to="/registration">
              <i class="material-icons black-text text-black">person</i>
            </RouterLink></li>
          </ul>
        </div>
      </nav>
    </div>

    <RouterView />
</template>

