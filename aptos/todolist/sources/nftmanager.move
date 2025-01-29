module devaddr::NftManager{
    use std::string;
    use std::signer;
    use std::string::String;
    use aptos_std::table;
    use aptos_std::table::Table;
    #[test_only]
    use aptos_framework::account::create_signer_for_test;
    #[test_only]
    use aptos_framework::create_signer;

    const CONTRACT_ADDRESS : address = @0x6;
    const ERROR_ADMIN_REQUIRED: u64 = 0;

    struct NFTTrait has store {
        color: String,
        speed: u64,
    }

    struct NFT has store{
        name: String,
        description: String,
        traits : vector<NFTTrait>,
        media: String
    }

    struct NFTCollection has key {
        nfts: Table<u64, NFT>,
        counter: u64
    }

    public entry fun create_nft(
        admin: &signer
    ){
        let admin_addr = signer::address_of(admin);
        assert!(admin_addr == CONTRACT_ADDRESS, ERROR_ADMIN_REQUIRED);

        if (!exists<NFTCollection>(admin_addr)){
            move_to(admin,  NFTCollection{
                nfts: table::new<u64, NFT>(),
                counter : 0
            })
        }
    }

    #[test]
    fun test_create_nft(){
        create_nft(&create_signer_for_test(CONTRACT_ADDRESS));
    }
}