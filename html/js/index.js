$(function(){
  // tab切换
  $('.tabTitle > li').each(function(i,item){
    $(this).on('click',function(){
      $(this).addClass('active').siblings('li').removeClass('active')
      $('.tabCon > li').eq(i).addClass('active').siblings('li').removeClass('active')
    })
  })

  // $.ajax({
  //   url:"http://localhost:8080/status",
  //   type:'get',
  //   success:function(data){
  //     console.log(data)
  //   }
  // })

  axios.get('http://localhost:8080/testdata/eth').then((response)=>{
    var data = response.data.data
    var count = data.peers.length
    var list = $('.tabCon>li').eq(0).find('.tabCon_r')
    if(response.data.code == 0){
      if(data.nodetype == 'multi'){
        $('.status_left').text('多机版')
      }else if(data.nodetype == 'single'){
        $('.status_left').text('单机版')
      }
      list.each(function(i,item){
        if(i<count){
          item.classList.add('run')
        }
      })
    }
  }).catch((err)=>{
    console.log(err)
  })

  $('#eth').on('click',function(){
    axios.get('http://localhost:8080/testdata/eth').then((response)=>{
      var data = response.data.data
      var count = data.peers.length
      var list = $('.tabCon>li').eq(0).find('.tabCon_r')
      console.log(list)
      if(response.data.code == 0){
        if(data.nodetype == 'multi'){
          $('.status_left').text('多机版')
        }else if(data.nodetype == 'single'){
          $('.status_left').text('单机版')
        }
        console.log(data)
        list.each(function(i,item){
          if(i<count){
            item.classList.add('run')
          }
        })
      }
    }).catch((err)=>{
      console.log(err)
    })
  })

  $('#eos').on('click',function(){
    axios.get('http://localhost:8080/testdata/eos').then((response)=>{
      var data = response.data.data
      var count = data.peers.length
      var list = $('.tabCon>li').eq(1).find('.tabCon_r')
      console.log(list)
      if(response.data.code == 0){
        if(data.nodetype == 'multi'){
          $('.status_left').text('多机版')
        }else if(data.nodetype == 'single'){
          $('.status_left').text('单机版')
        }
        console.log(data)
        list.each(function(i,item){
          if(i<count){
            item.classList.add('run')
          }
        })
      }
    }).catch((err)=>{
      console.log(err)
    })
  })

  $('#fisco').on('click',function(){
    axios.get('http://localhost:8080/testdata/fisco').then((response)=>{
      var data = response.data.data
      var count = data.peers.length
      var list = $('.tabCon>li').eq(2).find('.tabCon_r')
      console.log(list)
      if(response.data.code == 0){
        if(data.nodetype == 'multi'){
          $('.status_left').text('多机版')
        }else if(data.nodetype == 'single'){
          $('.status_left').text('单机版')
        }
        console.log(data)
        list.each(function(i,item){
          if(i<count){
            item.classList.add('run')
          }
        })
      }
    }).catch((err)=>{
      console.log(err)
    })
  })
})