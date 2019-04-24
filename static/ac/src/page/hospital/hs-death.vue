<template>
  <div class="wrapper">
    <!-- 头部 -->
    <headernav></headernav>
    <div class="main-content-wrap">
      <img src="../../assets/images/gone/pexels-photo-1930851.jpg" alt="">
      <el-row :gutter="20" style="padding: 0.2rem 1rem; background: #e8e8e8">
        <el-col :span="6"><img src="../../assets/images/gone/new1.jpg" alt="">宠物公墓</el-col>
        <el-col :span="6"><img src="../../assets/images/gone/new1.jpg" alt="">宠物墓地</el-col>
        <el-col :span="6"><img src="../../assets/images/gone/new1.jpg" alt="">宠物安乐死</el-col>
        <el-col :span="6"><img src="../../assets/images/gone/new1.jpg" alt="">宠物火化</el-col>
      </el-row>
    </div>
    <!-- 底部 -->
    <foot></foot>
  </div>
</template>
<script>
import headernav from "../../components/headernav.vue";
import foot from "../../components/foot"
export default {
  name:"hs_death",
  components:{
      headernav,foot 
  },
  data(){
    return{
      
    }
  },

}
</script>
<style scoped>
.main-content-wrap{
  width: 100%;
  background: #e8e8e8
}
.main-content-wrap > img{
  width: 100%;
}
.el-col{
  border: 1px solid #ffffff;
  background:#ffffff
}
.el-col img{
  width:2.5rem;
  height: 2rem;
}
</style>
