<!--
 Copyright 2022 Google LLC

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

      http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
-->

<template>
  <div
    class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 place-items-center gap-6 w-full"
  >
    <div class="grid place-items-end w-full">
      <img
        alt="Shifter"
        class="logo"
        src="@/assets/logo-shifter.svg"
        width="125"
        height="125"
      />
    </div>
    <div class="grid place-items-start inline-block align-middle w-full">
      <div class="">
        <p
          class="text-shifter-red text-4xl font-bold ease-in duration-400 hover:text-shifter-red"
        >
          Shifter
        </p>
        <p class="text-xl">Dang! We are sorry, something broke.</p>
      </div>
    </div>
  </div>
</template>

<style scoped>
.item {
  margin-top: 0;
  padding: 0.4rem 0 1rem calc(var(--section-gap) / 2);
}

i {
  top: calc(50% - 25px);
  left: -26px;
  position: absolute;
  border: 1px solid var(--color-border);
  background: var(--color-background);
  border-radius: 8px;
  width: 50px;
  height: 50px;
}

.item:before {
  content: " ";
  border-left: 1px solid var(--color-border);
  position: absolute;
  left: 0;
  bottom: calc(50% + 25px);
  height: calc(50% - 25px);
}

.item:after {
  content: " ";
  border-left: 1px solid var(--color-border);
  position: absolute;
  left: 0;
  top: calc(50% + 25px);
  height: calc(50% - 25px);
}

.item:first-of-type:before {
  display: none;
}

.item:last-of-type:after {
  display: none;
}

.item {
  margin-top: 2rem;
  display: flex;
}

.details {
  flex: 1;
  margin-left: 1rem;
}

i {
  display: flex;
  place-items: center;
  place-content: center;
  width: 32px;
  height: 32px;

  color: var(--color-text);
}

h3 {
  font-size: 1.2rem;
  font-weight: 500;
  margin-bottom: 0.4rem;
  color: var(--color-heading);
}
</style>
