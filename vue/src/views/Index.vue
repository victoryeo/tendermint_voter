<template>
  <div>
    <div class="sp-container">
      <sp-sign-in />
      <sp-bank-balances />
      <sp-token-send />
      <!-- this line is used by starport scaffolding # 4 -->
      <SpH3>
       Polls.Again.
      </SpH3>
      <poll-form />
    </div>
  </div>
</template>

<script>
import * as sp from "@tendermint/vue";
import PollForm from "../components/PollForm";

export default {
  components: { PollForm, ...sp },
};
</script>
