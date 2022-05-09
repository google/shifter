<template>
  <div>
    <div class="bg-shifter-black-soft">
      <nav
        class="
          container
          px-6
          py-4
          mx-auto
          md:flex md:justify-between md:items-center
        "
      >
        <div class="flex items-center justify-between">
           <router-link to="/"> 
            <img
                alt="Shifter"
                src="@/assets/logo-shifter.svg"
                class="mr-2"
                width="40"
                height="40"
                />
          </router-link>
          <router-link
            to="/"
            class="
              text-xl
              font-bold
              text-gray-800
              md:text-2xl
              hover:text-shifter-red
            "
            > Shifter
          </router-link>
          <!-- Mobile menu button -->
          <div @click="showMenu = !showMenu" class="flex md:hidden">
            <button
              type="button"
              class="
                text-gray-800
                hover:text-gray-400
                focus:outline-none focus:text-gray-400
              "
            >
              <svg viewBox="0 0 24 24" class="w-6 h-6 fill-current">
                <path
                  fill-rule="evenodd"
                  d="M4 5h16a1 1 0 0 1 0 2H4a1 1 0 1 1 0-2zm0 6h16a1 1 0 0 1 0 2H4a1 1 0 0 1 0-2zm0 6h16a1 1 0 0 1 0 2H4a1 1 0 0 1 0-2z"
                ></path>
              </svg>
            </button>
          </div>
        </div>


        <!-- Mobile Menu open: "block", Menu closed: "hidden" -->
        <ul
          :class="showMenu ? 'flex' : 'hidden'"
          class="
            flex-col
            mt-8
            space-y-4
            md:flex
            md:space-y-0
            md:flex-row
            md:items-center
            md:space-x-10
            md:mt-0
          "
        >
          <router-link to="/">
            <li class="text-sm font-bold hover:text-shifter-red-muted"> 
            Home
            </li>
          </router-link>
          <router-link to="/about">
            <li class="text-sm font-bold hover:text-shifter-red-muted"> 
            About
            </li>
          </router-link>
          <a target="_blank" href="https://github.com/google/shifter">
            <li class="text-sm font-bold hover:text-shifter-red-muted"> 
            Source
            </li>
          </a>
          <router-link to="/configure">
            <li class="text-sm font-bold hover:text-shifter-red-muted"> 
            Configure
            </li>
          </router-link>
          <router-link to="/convert">
            <li class="text-sm font-bold hover:text-shifter-red-muted"> 
            Convert
            </li>
          </router-link>
        </ul>
      </nav>
    </div>
  </div>
</template>
<script>
export default {
  data() {
    return {
      showMenu: false,
    };
  },
};
</script>