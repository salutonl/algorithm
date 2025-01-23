# -*- coding: utf-8 -*-
"""
@Time ： 2025/1/10 13:43
@Auth ： zhengtiantian
@File ：deepseek_api.py
@Description：
"""
# Please install OpenAI SDK first: `pip3 install openai`

from openai import OpenAI



client = OpenAI(api_key="sk-82853470ec74456cb41a373c8036d427", base_url="https://api.deepseek.com")
def ask_deepseek(quiz):

    response = client.chat.completions.create(
        model="deepseek-chat",
        messages=[
            {"role": "system", "content": "You are a helpful assistant"},
            {"role": "user", "content": "{}".format(quiz)},
        ],
        stream=False
    )
    return response.choices[0].message.content

if __name__ == '__main__':
    quiz = "I have a cloud server. How can I write GOlang code in real-time using the Goland IDE on my local server and synchronize it to the cloud server?"
    answer = ask_deepseek(quiz)
    print(answer)
