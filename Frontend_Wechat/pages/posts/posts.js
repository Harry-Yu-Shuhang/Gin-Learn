// var postData = require("../../data/data")
import {postList} from '../../data/data'

Page({

  /**
   * 页面的初始数据
   */
  data: {
  },

  /**
   * 生命周期函数--监听页面加载,也叫钩子函数(hook function)
   */
  onLoad(options) {
    //content已经是js对象了，所以不需要加花括号.this.setData要传入js对象objective
    //这里用于读取json数据
    this.setData({
      postList
    }) 
  },

  onGoToDetail(event){
    const pid = event.currentTarget.dataset.postId
    this.onUpdateViews(pid)
    wx.navigateTo({
      url: '/pages/post-detail/post-detail?pid='+pid,
    })
  },

  /*点击后浏览量+1 */
  // 点击后更新views并发送请求
  onUpdateViews(pid) {
    const post = this.data.postList.find(p => p.id === pid);
    
    // 判断post是否为undefined
    if (post === undefined) {
      console.log('pid不存在');  // 如果未找到文章，输出 "pid不存在"
      // 你还可以根据需要返回错误信息或者执行其他操作
      return
    } 

    // 发送请求到后端更新 views
    wx.request({
      url: 'http://localhost:8080/posts/click', // 替换为你的后端API地址
      method: 'POST',
      data: {
        id: pid,
        currentViews: post.views
      },
      success(res) {
        if (res.statusCode === 200 && res.data && res.data.newViews !== undefined) {// 如果后端返回成功，更新 views
          //遍历当前页面的数据列表 postList，对其中的每个文章对象 p 进行处理。
          const newPostList = this.data.postList.map(p => {
            if (p.id === pid) {
              return { ...p, views: res.data.newViews };
            }
            return p;
          });
          
          // 更新前端的 postList
          this.setData({
            postList: newPostList
          });
        }
      },
      fail(err) {
        console.error('获取后端请求失败', err);
      }
    });
  },

  /**
   * 生命周期函数--监听页面初次渲染完成
   */
  onReady() {

  },

  /**
   * 生命周期函数--监听页面显示
   */
  onShow() {

  },

  /**
   * 生命周期函数--监听页面隐藏
   */
  onHide() {

  },

  /**
   * 生命周期函数--监听页面卸载
   */
  onUnload() {

  },

  /**
   * 页面相关事件处理函数--监听用户下拉动作
   */
  onPullDownRefresh() {

  },

  /**
   * 页面上拉触底事件的处理函数
   */
  onReachBottom() {

  },

  /**
   * 用户点击右上角分享
   */
  onShareAppMessage() {

  }
})