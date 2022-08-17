import json, os

FOLDER_NAME="gentxs_from_strings"
GENTXS = [
    # testdedi
    """{"body":{"messages":[{"@type":"/cosmos.staking.v1beta1.MsgCreateValidator","description":{"moniker":"testdedi","identity":"","website":"https://google.com","security_contact":"dedi@email.io","details":""},"commission":{"rate":"0.069000000000000000","max_rate":"1.000000000000000000","max_change_rate":"0.250000000000000000"},"min_self_delegation":"1","delegator_address":"craft1r8sd9cgfnwzyt0dm06hnajy9mkljhl6suczml5","validator_address":"craftvaloper1r8sd9cgfnwzyt0dm06hnajy9mkljhl6s8uqjkt","pubkey":{"@type":"/cosmos.crypto.ed25519.PubKey","key":"n6LFj8THcTLuUPc5Hez06hqe2c/3D98sWxLMKAi3fJE="},"value":{"denom":"uexp","amount":"1000000"}}],"memo":"9014c44820d212e3aa84226d624465e668c5ef6a@95.217.113.126:26656","timeout_height":"0","extension_options":[],"non_critical_extension_options":[]},"auth_info":{"signer_infos":[{"public_key":{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"AjNfUYKrWXhVuv5d7isaPR1iVckVFap7Or2CcZnj3EP1"},"mode_info":{"single":{"mode":"SIGN_MODE_DIRECT"}},"sequence":"0"}],"fee":{"amount":[],"gas_limit":"200000","payer":"","granter":""},"tip":null},"signatures":["JFZ5nwQTR1akkUigNeFSkiKhmeT1lTH2SIy7L/74qE5nJ2NcR6LebDJFWrEx2MLLmAAba6UJ6O3mbtTJ/YfxDA=="]}""",
    # craft-validator machine
    """{"body":{"messages":[{"@type":"/cosmos.staking.v1beta1.MsgCreateValidator","description":{"moniker":"craft-validator","identity":"","website":"https://reece.sh","security_contact":"validator@email.io","details":""},"commission":{"rate":"0.050000000000000000","max_rate":"0.200000000000000000","max_change_rate":"0.050000000000000000"},"min_self_delegation":"1","delegator_address":"craft13vhr3gkme8hqvfyxd4zkmf5gaus840j5hwuqkh","validator_address":"craftvaloper13vhr3gkme8hqvfyxd4zkmf5gaus840j5v27flg","pubkey":{"@type":"/cosmos.crypto.ed25519.PubKey","key":"fkQLjn1XsAvmv+jwqjgCNJHW4yGa4SzuUaAfUg7Dxvk="},"value":{"denom":"uexp","amount":"1000000"}}],"memo":"47c82ac6c012cccdc2e0a3a3372e62078066cc51@65.109.38.251:26656","timeout_height":"0","extension_options":[],"non_critical_extension_options":[]},"auth_info":{"signer_infos":[{"public_key":{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"AwS2eOmDNypUhlwWCj/b7JwiszaY/YiUMZa9E1xFX3Gf"},"mode_info":{"single":{"mode":"SIGN_MODE_DIRECT"}},"sequence":"0"}],"fee":{"amount":[],"gas_limit":"200000","payer":"","granter":""},"tip":null},"signatures":["QjAGq1GjswgrWqeDCsO5aaquA6t52nQ50BwSdjtb8s0NVTvEayzDxqrz8otl6FZ5yQA/do7NIW6MDdYE+dgW+w=="]}"""
]

os.makedirs(FOLDER_NAME, exist_ok=True)

for gentx in GENTXS:
    v = json.loads(gentx)
    moniker = v['body']['messages'][0]['description']['moniker']
    memo = v['body']['memo']
    if '@' in memo: print(f"{moniker} - {memo}")

    # create file named moniker.json
    with open(f"{FOLDER_NAME}/{moniker}.json", "w") as f:
        json.dump(v, f)
        # print(f"{moniker}.json created")

print(f"ALL GENTXS CREATED FROM STRINGS. MOVE {FOLDER_NAME} INTO 'gentx' folder. Then run add-genesis-accounts.py")