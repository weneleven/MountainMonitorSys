import json

with open('data.json', 'r') as file:
    raw_data = json.load(file)


new_data = []

for item in raw_data:
    new_item = {
        "sensor_name": item["点名"],
        "collect_time": item["采集时间"],
        "value_x": float(item["X坐标值(m)"]),
        "value_y": float(item["Y坐标值(m)"]),
        "value_h": float(item["H坐标值"]),
        "sum_x": float(item["X方向累计(mm)"]),
        "sum_y": float(item["Y方向累计(mm)"]),
        "sum_h": float(item["H方向累计(mm)"]),
        "now_x": float(item["X方向本次(mm)"]),
        "now_y": float(item["Y方向本次(mm)"]),
        "now_h": float(item["H方向本次(mm)"])
    }
    original_time = new_item["collect_time"]
    parsed_time = original_time.replace(" ", "T") + "+08:00"
    new_item["collect_time"] = parsed_time
    new_data.append(new_item)

with open('new_data.json', 'w') as new_file:
    json.dump(new_data, new_file, indent=4)

print("转换完成，新的JSON数据已保存到new_data.json文件中。")
