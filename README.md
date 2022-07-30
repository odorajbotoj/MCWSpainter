# MCWSpainter
!!!有重大bug待修复!!!像素画绘制(针对1.17以上基岩版),与ws搭配使用
### 用法
+ 首先,注释是个好东西...
+ 你需要打开一个图片文件,并使用image.Decode来获取一个image.Image对象
+ 然后调用MapImage(image.Image, ColorSpace, int, time.Duration).其中image.Image为你解码的文件对象, ColorSpace为你的颜色空间(我们做了一个用于三维像素画的颜色空间DefaultColorSpace, 你也可以自定义一个), int为生成图像的基准y轴, time.Duration为延时, 用于设置每条命令执行间的延时, 防止太快的执行卡死客户端.
+ MapImage函数返回一个Painter. 接下来, 你需要不停调用其Next()方法, 并获取填充方块的X Y Z坐标与方块名 数据值.
+ Next()中已经有延时了. 但是它不会在绘制完成后报错,反而会从头重新填充. 你可以使用GetSize()方法来获取需要循环的次数.
+ 其他的...自己看demo吧
### 存在的问题
+ 听说3D地图画(不同高度的方块在地图上有不同的阴影)仅在Java版有效.
+ 所以如果你要导入普通的平面地图画, 无需理会方法回传的Y轴参数, 直接按照输入的Y轴来!
### 外部调用
+ 详见我的另一个项目wetsponge的demo
##### Have a nice day! ;P
