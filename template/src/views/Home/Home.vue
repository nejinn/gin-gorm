<template>
  <nly-container fluid>
    <nly-row>
      <nly-col
        xs="12"
        md="3"
        sm="6"
        v-for="(item, index) in infoBox"
        :key="index"
      >
        <nly-infobox :bg-variant="item.variant">
          <nly-infobox-icon :icon="item.icon" />
          <nly-infobox-body
            :number="item.num === 0 ? '0' : item.num"
            :text="item.name"
            :progress-value="item.growth_rate === 0 ? 1 : item.growth_rate"
            :progress-description="
              '最近' + item.day + '天增长' + item.growth_rate + '%'
            "
          />
        </nly-infobox>
      </nly-col>
    </nly-row>
    <nly-row>
      <nly-col
        xs="12"
        md="3"
        sm="6"
        v-for="(item, index) in infoBox1"
        :key="index"
      >
        <nly-infobox>
          <nly-infobox-icon :icon="item.icon" :bg-variant="item.variant" />
          <nly-infobox-body
            :number="item.num === 0 ? '0' : item.num"
            :text="item.name"
          />
        </nly-infobox>
      </nly-col>
    </nly-row>
  </nly-container>
</template>

<script>
export default {
  data() {
    return {
      infoBox: undefined,
      infoBox1: undefined
    };
  },
  mounted() {
    const obj = this;
    obj.isBusy = true;
    const params = {
      url: "/dashboard"
    };
    this.$api.HttpsDashboard.getInfoBox(obj, params);
  }
};
</script>
