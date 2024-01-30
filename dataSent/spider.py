from bs4 import BeautifulSoup
import json
from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
from PIL import Image
import base64
from io import BytesIO
import tesserocr
import time
import re


def recognize_captcha():
    # 找到验证码图片元素
    captcha_element = driver.find_element(By.CSS_SELECTOR, ".login-code-img")
    captcha_src = captcha_element.get_attribute('src')

    # 解码base64图像
    header, encoded = captcha_src.split(",", 1)
    captcha_image = Image.open(BytesIO(base64.b64decode(encoded)))

    # 使用 tesserocr 识别验证码
    with tesserocr.PyTessBaseAPI(path=tessdata_path) as api:
        api.SetImage(captcha_image)
        captcha_code = api.GetUTF8Text().strip()

    # 清除非字母数字字符，截取前四位，如果为空设置为 '1145'
    captcha_code = re.sub(r'[^A-Za-z0-9]', '', captcha_code)[:4] or '1145'
    return captcha_code


def login(captcha_code):
    # 等待验证码输入框出现
    wait = WebDriverWait(driver, 10)
    captcha_input = wait.until(EC.presence_of_element_located((By.XPATH, "//input[@placeholder='验证码']")))

    # 先清除输入框中的现有内容，然后输入新的验证码
    captcha_input.clear()
    captcha_input.send_keys(captcha_code)

    # 点击登录按钮
    login_button = driver.find_element(By.CSS_SELECTOR, ".el-button--primary")
    login_button.click()

    # 等待页面加载
    time.sleep(2)

    # 获取登录后的页面代码
    return driver.page_source


# 创建一个浏览器实例
driver = webdriver.Chrome()

tessdata_path = 'E:/tesserocr/Tesseract-OCR/tessdata'  # 替换为实际 tessdata 文件夹所在的路径

# 设置目标网址
url = "http://47.96.105.25:4181/realtime/pointdatalist"

# 使用浏览器打开网页
driver.get(url)

# 添加Cookie
cookies = [
    {'name': 'TDTSESID', 'value': '53b4229e8ec38ce2514f781e5af07988|8a99c1bdeb67a53f52230a32ac84b973'},
    {'name': 'username', 'value': 'sxlfgxpdj'},
    {'name': 'password',
     'value': 'JQXcXHXaBHMCdMg0K0vYARZUgave44RaY8PD5K7Mbq4HRsJ/H4Z9dCMIuPLCP7+RfSIPrJwmYshgW6GhqymQRg=='},
    {'name': 'rememberMe', 'value': 'true'}
]

for cookie in cookies:
    driver.add_cookie(cookie)

# 刷新页面以应用Cookie
driver.refresh()

# 等待页面加载
time.sleep(5)

# 尝试登录，如果验证码错误则重试
max_attempts = 50  # 设置最大尝试次数
flag = 0
for attempt in range(max_attempts):
    captcha_code = recognize_captcha()
    page_source = login(captcha_code)

    if "验证码错误" not in page_source:
        print("正确")
        flag = 1
        break
    else:
        print("验证码错误，正在重试...")


if flag == 1:
    # 等待页面加载
    driver.get("http://47.96.105.25:4181/realtime/pointdatalist")
    # 等待页面加载
    time.sleep(5)
    # 返回新页面的内容
    html_content = driver.page_source
    soup = BeautifulSoup(html_content, 'lxml')

    # 提取表格行
    rows = soup.find_all('tr', class_='el-table__row')

    # 解析数据
    data = []

    for row in rows:
        cells = row.find_all('td')
        if len(cells) >= 12:  # 确保有足够的单元格
            # 提取并转换数据
            sensor_name = cells[3].get_text().strip()
            collect_time = cells[2].get_text().strip()
            value_x = float(cells[4].get_text().strip())
            value_y = float(cells[5].get_text().strip())
            value_h = float(cells[6].get_text().strip())
            sum_x = float(cells[7].get_text().strip())
            sum_y = float(cells[8].get_text().strip())
            sum_h = float(cells[9].get_text().strip())
            now_x = float(cells[10].get_text().strip())
            now_y = float(cells[11].get_text().strip())
            now_h = float(cells[12].get_text().strip())

            # 创建字典并添加到列表
            data.append({
                "sensor_name": sensor_name,
                "collect_time": collect_time,
                "value_x": value_x,
                "value_y": value_y,
                "value_h": value_h,
                "sum_x": sum_x,
                "sum_y": sum_y,
                "sum_h": sum_h,
                "now_x": now_x,
                "now_y": now_y,
                "now_h": now_h
            })

    # 将数据转换为 JSON
    json_data = json.dumps(data, indent=4)
    print(json_data)

# 关闭浏览器
driver.quit()
