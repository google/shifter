/**
 * Copyright 2022 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

module.exports = {
  darkMode: 'class',
  content: ["./index.html", "./src/**/*.{vue,js,ts,jsx,tsx}"],
  theme: {
    colors: {
      "shifter-black": "#121212",
      "shifter-black-soft": "#282828",
      "shifter-black-mute": "#121212",

      "shifter-blur": "#121212",

      "shifter-grey-background": "#fafafa",

      "shifter-white": "#ffffff",
      "shifter-white-soft": "#e2e2e2",
      "shifter-white-mute": "#f2f2f2",

      "shifter-red": "#db212d",
      "shifter-red-soft": "#bd1c27",
      "shifter-red-muted": "#a61e27",

      "shifter-blue": "#6200EE",

      "dummy-green": "#4cbd46",
      "dummy-blue": "#6995cf",
    },

    extend: {},
  },
  plugins: [],
};
